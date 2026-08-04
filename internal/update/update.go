package update

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/faman-project/faman/internal/parser"
	"github.com/faman-project/faman/internal/renderer"
)

const (
	repoOwner  = "erfankasraie"
	repoName   = "Faman"
	apiBase    = "https://api.github.com/repos/" + repoOwner + "/" + repoName
	archiveURL = "https://github.com/" + repoOwner + "/" + repoName + "/archive/refs/heads/main.tar.gz"
)

// Options controls update behavior.
type Options struct {
	CurrentVersion string
	CheckOnly      bool
	PagesOnly      bool
	Force          bool
	Verify         bool // with --pages: use Release archive + SHA256SUMS
}

// Run executes the update command.
func Run(opt Options) error {
	useColor := renderer.ColorEnabled()
	printTitle(useColor)

	if opt.PagesOnly {
		return refreshPages(opt, useColor)
	}

	remote, src, err := latestRemoteVersion()
	if err != nil {
		printWarn(useColor, "نتوانستیم GitHub را بخوانیم: "+err.Error())
		printOfflineHints(useColor)
		if opt.CheckOnly {
			return err
		}
		fmt.Println()
		return nil
	}

	cur := strings.TrimPrefix(strings.TrimSpace(opt.CurrentVersion), "v")
	rem := strings.TrimPrefix(strings.TrimSpace(remote), "v")

	printKV(useColor, "نسخه محلی", cur)
	printKV(useColor, "آخرین روی GitHub", rem+"  ("+src+")")

	switch compareVersions(cur, rem) {
	case 0:
		printOK(useColor, "با آخرین انتشار هم‌تراز هستید.")
	case -1:
		printWarn(useColor, "نسخهٔ جدیدتری روی GitHub هست.")
	default:
		printOK(useColor, "نسخهٔ محلی جدیدتر از آخرین انتشار برچسب‌خورده است (احتمالاً build از main).")
	}

	if opt.CheckOnly {
		return nil
	}

	fmt.Println()
	printSection(useColor, "به‌روزرسانی صفحات")
	fmt.Println(dim(useColor, "  faman update --pages"))
	fmt.Println(dim(useColor, "  faman update --pages --verify   # + SHA256 از Release"))
	fmt.Println()
	printSection(useColor, "به‌روزرسانی باینری")
	if runtime.GOOS == "windows" {
		fmt.Println(dim(useColor, "  irm https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/install.ps1 | iex"))
	} else {
		fmt.Println(dim(useColor, "  curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/get.sh | bash"))
	}
	fmt.Println(dim(useColor, "  تأیید دستی: sha256sum -c SHA256SUMS"))
	fmt.Println()
	printSection(useColor, "لینک‌ها")
	fmt.Println(dim(useColor, "  https://github.com/erfankasraie/Faman/releases"))
	return nil
}

func refreshPages(opt Options, useColor bool) error {
	dest, err := pagesInstallDir(opt.Force)
	if err != nil {
		return err
	}
	printKV(useColor, "مقصد صفحات", dest)

	if opt.Verify {
		printOK(useColor, "دانلود از Release + تأیید SHA256SUMS…")
		tag, name, sha, err := downloadPagesFromRelease(dest)
		if err != nil {
			return err
		}
		printKV(useColor, "تگ Release", tag)
		printKV(useColor, "آرشیو", name)
		printOK(useColor, "SHA256 تأیید شد: "+sha)
	} else {
		printOK(useColor, "دانلود آرشیو شاخه main…")
		sha, err := downloadPagesFromMain(dest)
		if err != nil {
			return err
		}
		printKV(useColor, "SHA256 آرشیو", sha)
		printWarn(useColor, "بدون --verify فقط هش چاپ می‌شود (main هر commit عوض می‌شود).")
	}

	n := countMD(dest)
	printOK(useColor, fmt.Sprintf("صفحات به‌روز شد (%d فایل .md).", n))
	fmt.Println(dim(useColor, "  export FAMAN_PAGES=\""+dest+"\""))
	return nil
}

func latestRemoteVersion() (tag, source string, err error) {
	client := &http.Client{Timeout: 12 * time.Second}

	req, _ := http.NewRequest(http.MethodGet, apiBase+"/releases/latest", nil)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("User-Agent", "faman-update")
	resp, err := client.Do(req)
	if err == nil {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			var body struct {
				TagName string `json:"tag_name"`
			}
			if json.NewDecoder(resp.Body).Decode(&body) == nil && body.TagName != "" {
				return body.TagName, "release", nil
			}
		}
	}

	req2, _ := http.NewRequest(http.MethodGet, apiBase+"/tags?per_page=5", nil)
	req2.Header.Set("Accept", "application/vnd.github+json")
	req2.Header.Set("User-Agent", "faman-update")
	resp2, err := client.Do(req2)
	if err != nil {
		return "", "", err
	}
	defer resp2.Body.Close()
	if resp2.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("GitHub API status %d", resp2.StatusCode)
	}
	var tags []struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(resp2.Body).Decode(&tags); err != nil {
		return "", "", err
	}
	if len(tags) == 0 {
		return "main", "branch", nil
	}
	return tags[0].Name, "tag", nil
}

func pagesInstallDir(force bool) (string, error) {
	if env := os.Getenv("FAMAN_PAGES"); env != "" {
		if writableDir(env) || force {
			return env, nil
		}
	}
	if cur, err := parser.PagesDir(); err == nil {
		if underHome(cur) && (writableDir(cur) || force) {
			return cur, nil
		}
		if force && writableDir(filepath.Dir(cur)) {
			return cur, nil
		}
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	if runtime.GOOS == "windows" {
		base := os.Getenv("LOCALAPPDATA")
		if base == "" {
			base = filepath.Join(home, "AppData", "Local")
		}
		return filepath.Join(base, "faman", "pages", "fa"), nil
	}
	return filepath.Join(home, ".local", "share", "faman", "pages", "fa"), nil
}

func underHome(path string) bool {
	home, err := os.UserHomeDir()
	if err != nil {
		return false
	}
	abs, err := filepath.Abs(path)
	if err != nil {
		return false
	}
	rel, err := filepath.Rel(home, abs)
	if err != nil {
		return false
	}
	return !strings.HasPrefix(rel, "..")
}

func writableDir(dir string) bool {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return false
	}
	f, err := os.CreateTemp(dir, ".faman-write-*")
	if err != nil {
		return false
	}
	name := f.Name()
	_ = f.Close()
	_ = os.Remove(name)
	return true
}

func countMD(dir string) int {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 0
	}
	n := 0
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".md") {
			n++
		}
	}
	return n
}

func compareVersions(a, b string) int {
	a = strings.TrimPrefix(strings.ToLower(a), "v")
	b = strings.TrimPrefix(strings.ToLower(b), "v")
	if a == b {
		return 0
	}
	abase := strings.Split(a, "-")[0]
	bbase := strings.Split(b, "-")[0]
	ap := strings.Split(abase, ".")
	bp := strings.Split(bbase, ".")
	for len(ap) < 3 {
		ap = append(ap, "0")
	}
	for len(bp) < 3 {
		bp = append(bp, "0")
	}
	for i := 0; i < 3; i++ {
		ai, bi := atoi(ap[i]), atoi(bp[i])
		if ai < bi {
			return -1
		}
		if ai > bi {
			return 1
		}
	}
	hasPreA := strings.Contains(a, "-")
	hasPreB := strings.Contains(b, "-")
	if hasPreA && !hasPreB {
		return -1
	}
	if !hasPreA && hasPreB {
		return 1
	}
	return 0
}

func atoi(s string) int {
	n := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			break
		}
		n = n*10 + int(r-'0')
	}
	return n
}

func printTitle(useColor bool) {
	t := "∧  به‌روزرسانی faman"
	if useColor {
		fmt.Println(renderer.TitleStyle.Render(t))
		fmt.Println(renderer.DimStyle.Render(strings.Repeat("─", 36)))
	} else {
		fmt.Println(t)
		fmt.Println(strings.Repeat("-", 36))
	}
	fmt.Println()
}

func printKV(useColor bool, k, v string) {
	if useColor {
		fmt.Printf("  %s  %s\n", renderer.MetaStyle.Render(k+":"), v)
	} else {
		fmt.Printf("  %s: %s\n", k, v)
	}
}

func printOK(useColor bool, msg string) {
	if useColor {
		fmt.Println(renderer.SuccessStyle.Render("  ✓  " + msg))
	} else {
		fmt.Println("  OK  " + msg)
	}
}

func printWarn(useColor bool, msg string) {
	if useColor {
		fmt.Println(renderer.WarningStyle.Render("  !  " + msg))
	} else {
		fmt.Println("  !  " + msg)
	}
}

func printSection(useColor bool, title string) {
	if useColor {
		fmt.Println(renderer.SectionStyle.Render(title))
	} else {
		fmt.Println(title)
	}
}

func dim(useColor bool, s string) string {
	if useColor {
		return renderer.DimStyle.Render(s)
	}
	return s
}

func printOfflineHints(useColor bool) {
	fmt.Println()
	printSection(useColor, "آفلاین / دستی")
	fmt.Println(dim(useColor, "  curl -fsSL https://raw.githubusercontent.com/erfankasraie/Faman/main/scripts/get.sh | bash"))
	fmt.Println(dim(useColor, "  faman update --pages --verify"))
}
