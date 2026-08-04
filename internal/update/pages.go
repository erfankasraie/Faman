package update

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// downloadPagesFromMain downloads GitHub main branch archive, prints SHA256, extracts pages/fa.
func downloadPagesFromMain(dest string) (sha string, err error) {
	tmpFile, err := os.CreateTemp("", "faman-main-*.tar.gz")
	if err != nil {
		return "", err
	}
	path := tmpFile.Name()
	_ = tmpFile.Close()
	defer os.Remove(path)

	if err := downloadToFile(archiveURL, path); err != nil {
		return "", err
	}
	sha, err = fileSHA256(path)
	if err != nil {
		return "", err
	}
	if err := extractPagesFaFromTarGz(path, dest); err != nil {
		return sha, err
	}
	return sha, nil
}

// downloadPagesFromRelease downloads platform archive from latest GitHub Release,
// verifies SHA256 against release SHA256SUMS, extracts pages/fa.
func downloadPagesFromRelease(dest string) (tag, archiveName, sha string, err error) {
	tag, _, err = latestRemoteVersion()
	if err != nil {
		return "", "", "", err
	}
	tag = strings.TrimSpace(tag)
	ver := strings.TrimPrefix(tag, "v")

	archiveName = releaseArchiveName(ver)
	base := fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/", repoOwner, repoName, tag)
	sumsURL := base + "SHA256SUMS"
	archURL := base + archiveName

	tmpDir, err := os.MkdirTemp("", "faman-rel-*")
	if err != nil {
		return tag, archiveName, "", err
	}
	defer os.RemoveAll(tmpDir)

	sumsPath := filepath.Join(tmpDir, "SHA256SUMS")
	archPath := filepath.Join(tmpDir, archiveName)
	if err := downloadToFile(sumsURL, sumsPath); err != nil {
		return tag, archiveName, "", fmt.Errorf("SHA256SUMS: %w (آیا Release آرتیفکت دارد؟)", err)
	}
	if err := downloadToFile(archURL, archPath); err != nil {
		return tag, archiveName, "", fmt.Errorf("archive: %w", err)
	}

	f, err := os.Open(sumsPath)
	if err != nil {
		return tag, archiveName, "", err
	}
	sums, err := parseSHA256SUMS(f)
	_ = f.Close()
	if err != nil {
		return tag, archiveName, "", err
	}
	if err := verifyFileAgainstSUMS(archPath, sums); err != nil {
		return tag, archiveName, "", err
	}
	sha, err = fileSHA256(archPath)
	if err != nil {
		return tag, archiveName, "", err
	}

	switch {
	case strings.HasSuffix(archiveName, ".tar.gz"):
		err = extractPagesFaFromTarGz(archPath, dest)
	case strings.HasSuffix(archiveName, ".zip"):
		err = extractPagesFaFromZip(archPath, dest)
	default:
		err = fmt.Errorf("unknown archive type: %s", archiveName)
	}
	return tag, archiveName, sha, err
}

func releaseArchiveName(version string) string {
	switch runtime.GOOS {
	case "windows":
		return fmt.Sprintf("faman-%s-windows-amd64.zip", version)
	case "darwin":
		if runtime.GOARCH == "arm64" {
			return fmt.Sprintf("faman-%s-darwin-arm64.tar.gz", version)
		}
		return fmt.Sprintf("faman-%s-darwin-amd64.tar.gz", version)
	default: // linux and others
		if runtime.GOARCH == "arm64" {
			return fmt.Sprintf("faman-%s-linux-arm64.tar.gz", version)
		}
		return fmt.Sprintf("faman-%s-linux-amd64.tar.gz", version)
	}
}

func extractPagesFaFromTarGz(archivePath, dest string) error {
	f, err := os.Open(archivePath)
	if err != nil {
		return err
	}
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer gz.Close()
	return extractPagesFaFromTar(tar.NewReader(gz), dest)
}

func extractPagesFaFromTar(tr *tar.Reader, dest string) error {
	tmp, err := os.MkdirTemp("", "faman-pages-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp)

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
			// release layout: faman-VERSION-os/pages/fa/...
			continue
		}
		rel := name[idx+len("/pages/fa/"):]
		if rel == "" || strings.Contains(rel, "..") {
			continue
		}
		target := filepath.Join(tmp, filepath.FromSlash(rel))
		switch hdr.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, 0o755); err != nil {
				return err
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
				return err
			}
			out, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o644)
			if err != nil {
				return err
			}
			if _, err := io.Copy(out, tr); err != nil {
				_ = out.Close()
				return err
			}
			_ = out.Close()
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
