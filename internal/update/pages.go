package update

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func downloadPagesTo(dest string) error {
	client := &http.Client{Timeout: 120 * time.Second}
	req, err := http.NewRequest(http.MethodGet, archiveURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "faman-update")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed: HTTP %d", resp.StatusCode)
	}

	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}
	defer gz.Close()

	tr := tar.NewReader(gz)
	tmp, err := os.MkdirTemp("", "faman-pages-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp)

	prefix := "" // e.g. Faman-main/pages/fa/
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		name := filepath.ToSlash(hdr.Name)
		idx := strings.Index(name, "/pages/fa/")
		if idx < 0 {
			// also allow exact pages/fa at start after repo root
			continue
		}
		rel := name[idx+len("/pages/fa/"):]
		if rel == "" || strings.HasSuffix(name, "/pages/fa") {
			continue
		}
		if strings.Contains(rel, "..") {
			continue
		}
		target := filepath.Join(tmp, filepath.FromSlash(rel))
		switch hdr.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, 0o755); err != nil {
				return err
			}
		case tar.TypeReg, tar.TypeRegA:
			if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
				return err
			}
			f, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
			if err != nil {
				return err
			}
			if _, err := io.Copy(f, tr); err != nil {
				_ = f.Close()
				return err
			}
			_ = f.Close()
			if prefix == "" {
				prefix = name[:idx+len("/pages/fa/")]
			}
		}
	}

	entries, err := os.ReadDir(tmp)
	if err != nil {
		return err
	}
	if len(entries) == 0 {
		return fmt.Errorf("no pages/fa found in archive")
	}

	if err := os.MkdirAll(dest, 0o755); err != nil {
		return err
	}
	// Replace contents: remove old md files then copy
	old, _ := os.ReadDir(dest)
	for _, e := range old {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".md") {
			_ = os.Remove(filepath.Join(dest, e.Name()))
		}
	}
	return copyTree(tmp, dest)
}

func copyTree(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if info.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()
		out, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
		if err != nil {
			return err
		}
		defer out.Close()
		_, err = io.Copy(out, in)
		return err
	})
}
