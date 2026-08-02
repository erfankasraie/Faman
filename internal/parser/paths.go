package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// pagesDir returns the directory containing markdown pages.
// Search order (cross-platform):
//  1. FAMAN_PAGES
//  2. Next to the executable (…/pages/fa)
//  3. OS user data dirs (Windows LOCALAPPDATA, Unix ~/.local/share, …)
//  4. Unix FHS share paths (linux/darwin only)
//  5. Current working directory (development)
func pagesDir() (string, error) {
	if env := os.Getenv("FAMAN_PAGES"); env != "" {
		if st, err := os.Stat(env); err == nil && st.IsDir() {
			return env, nil
		}
	}

	var candidates []string

	if exe, err := os.Executable(); err == nil {
		exeDir := filepath.Dir(exe)
		candidates = append(candidates,
			filepath.Join(exeDir, "pages", "fa"),
			filepath.Join(exeDir, "..", "pages", "fa"),
			filepath.Join(exeDir, "..", "share", "faman", "pages", "fa"),
		)
	}

	candidates = append(candidates, userPagesCandidates()...)

	if runtime.GOOS != "windows" {
		candidates = append(candidates,
			"/usr/local/share/faman/pages/fa",
			"/usr/share/faman/pages/fa",
			"/opt/faman/pages/fa",
		)
	}

	if cwd, err := os.Getwd(); err == nil {
		candidates = append(candidates,
			filepath.Join(cwd, "pages", "fa"),
			filepath.Join(cwd, "..", "pages", "fa"),
		)
	}

	for _, dir := range candidates {
		if dir == "" {
			continue
		}
		if st, err := os.Stat(dir); err == nil && st.IsDir() {
			return dir, nil
		}
	}

	hint := "set FAMAN_PAGES to your pages/fa folder"
	if runtime.GOOS == "windows" {
		hint = "set FAMAN_PAGES or copy pages\\fa to %LOCALAPPDATA%\\faman\\pages\\fa (see docs/windows.md)"
	} else {
		hint = "set FAMAN_PAGES or install pages to /usr/local/share/faman/pages/fa"
	}
	return "", fmt.Errorf("pages directory not found (%s)", hint)
}

func userPagesCandidates() []string {
	var out []string

	if runtime.GOOS == "windows" {
		if la := os.Getenv("LOCALAPPDATA"); la != "" {
			out = append(out, filepath.Join(la, "faman", "pages", "fa"))
		}
		if ad := os.Getenv("APPDATA"); ad != "" {
			out = append(out, filepath.Join(ad, "faman", "pages", "fa"))
		}
	}

	// XDG / cross-platform user config & home
	if cfg, err := os.UserConfigDir(); err == nil && cfg != "" {
		out = append(out, filepath.Join(cfg, "faman", "pages", "fa"))
	}
	if home, err := os.UserHomeDir(); err == nil && home != "" {
		out = append(out,
			filepath.Join(home, ".local", "share", "faman", "pages", "fa"),
			filepath.Join(home, "faman", "pages", "fa"),
		)
	}
	return out
}
