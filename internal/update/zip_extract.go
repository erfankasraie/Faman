package update

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func extractPagesFaFromZip(archivePath, dest string) error {
	r, err := zip.OpenReader(archivePath)
	if err != nil {
		return err
	}
	defer r.Close()

	tmp, err := os.MkdirTemp("", "faman-pages-zip-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp)

	found := 0
	for _, f := range r.File {
		name := filepath.ToSlash(f.Name)
		idx := strings.Index(name, "/pages/fa/")
		if idx < 0 {
			continue
		}
		rel := name[idx+len("/pages/fa/"):]
		if rel == "" || strings.Contains(rel, "..") {
			continue
		}
		target := filepath.Join(tmp, filepath.FromSlash(rel))
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(target, 0o755); err != nil {
				return err
			}
			continue
		}
		if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
			return err
		}
		rc, err := f.Open()
		if err != nil {
			return err
		}
		out, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
		if err != nil {
			_ = rc.Close()
			return err
		}
		_, err = io.Copy(out, rc)
		_ = out.Close()
		_ = rc.Close()
		if err != nil {
			return err
		}
		found++
	}
	if found == 0 {
		return fmt.Errorf("no pages/fa found in zip")
	}
	if err := os.MkdirAll(dest, 0o755); err != nil {
		return err
	}
	old, _ := os.ReadDir(dest)
	for _, e := range old {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".md") {
			_ = os.Remove(filepath.Join(dest, e.Name()))
		}
	}
	return copyTree(tmp, dest)
}
