package update

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func fileSHA256(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func downloadToFile(url, dest string) error {
	client := &http.Client{Timeout: 180 * time.Second}
	req, err := http.NewRequest(http.MethodGet, url, nil)
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
		return fmt.Errorf("download %s: HTTP %d", url, resp.StatusCode)
	}
	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	return err
}

// parseSHA256SUMS maps basename -> hex digest (GNU coreutils format).
func parseSHA256SUMS(r io.Reader) (map[string]string, error) {
	out := map[string]string{}
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		// "<hex>  <filename>" or "<hex> *filename"
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		hexSum := fields[0]
		name := fields[1]
		name = strings.TrimPrefix(name, "*")
		out[filepath.Base(name)] = strings.ToLower(hexSum)
	}
	return out, sc.Err()
}

func verifyFileAgainstSUMS(filePath string, sums map[string]string) error {
	base := filepath.Base(filePath)
	want, ok := sums[base]
	if !ok {
		return fmt.Errorf("فایل %s در SHA256SUMS نیست", base)
	}
	got, err := fileSHA256(filePath)
	if err != nil {
		return err
	}
	if !strings.EqualFold(got, want) {
		return fmt.Errorf("SHA256 mismatch for %s\n  expected: %s\n  got:      %s", base, want, got)
	}
	return nil
}
