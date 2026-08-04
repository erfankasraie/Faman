package app

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/cobra"

	"github.com/faman-project/faman/internal/parser"
	"github.com/faman-project/faman/internal/renderer"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "بررسی محیط اجرا (pages، locale، نسخه)",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runDoctor()
	},
}

func runDoctor() error {
	useColor := renderer.ColorEnabled()
	ok := 0
	warn := 0

	printDocTitle := func() {
		if useColor {
			fmt.Println(renderer.TitleStyle.Render("∧  faman doctor"))
			fmt.Println(renderer.DimStyle.Render(strings.Repeat("─", 40)))
		} else {
			fmt.Println("faman doctor")
		}
	}
	lineOK := func(label, detail string) {
		ok++
		if useColor {
			fmt.Printf("  %s  %s  %s\n", renderer.SuccessStyle.Render("✓"), label, renderer.DimStyle.Render(detail))
		} else {
			fmt.Printf("  OK   %s  %s\n", label, detail)
		}
	}
	lineWarn := func(label, detail string) {
		warn++
		if useColor {
			fmt.Printf("  %s  %s  %s\n", renderer.WarningStyle.Render("!"), label, renderer.DimStyle.Render(detail))
		} else {
			fmt.Printf("  WARN %s  %s\n", label, detail)
		}
	}
	lineFail := func(label, detail string) {
		warn++
		if useColor {
			fmt.Printf("  %s  %s  %s\n", renderer.ErrorStyle.Render("✗"), label, renderer.DimStyle.Render(detail))
		} else {
			fmt.Printf("  FAIL %s  %s\n", label, detail)
		}
	}

	printDocTitle()

	// Version / runtime
	lineOK("version", Version())
	lineOK("goos/goarch", runtime.GOOS+"/"+runtime.GOARCH)

	// Pages dir
	dir, err := parser.PagesDir()
	if err != nil {
		lineFail("pages", err.Error())
	} else {
		lineOK("pages", dir)
		n := 0
		if entries, e := os.ReadDir(dir); e == nil {
			for _, e := range entries {
				if !e.IsDir() && strings.HasSuffix(e.Name(), ".md") {
					n++
				}
			}
		}
		lineOK("page count", fmt.Sprintf("%d", n))
		// writable?
		if isWritableDir(dir) {
			lineOK("pages writable", "yes (update --pages ممکن است)")
		} else {
			lineWarn("pages writable", "no — update --pages به مسیر کاربر می‌رود")
		}
	}

	// FAMAN_PAGES
	if v := os.Getenv("FAMAN_PAGES"); v != "" {
		lineOK("FAMAN_PAGES", v)
	} else {
		lineOK("FAMAN_PAGES", "(unset — از مسیر پیش‌فرض)")
	}

	// Locale / UTF-8
	lang := os.Getenv("LANG")
	if lang == "" {
		lang = os.Getenv("LC_ALL")
	}
	if lang == "" {
		if runtime.GOOS == "windows" {
			lineWarn("locale", "LANG خالی (روی ویندوز رایج است)")
		} else {
			lineWarn("locale", "LANG/LC_ALL خالی — UTF-8 توصیه می‌شود")
		}
	} else if strings.Contains(strings.ToUpper(lang), "UTF-8") || strings.Contains(strings.ToUpper(lang), "UTF8") {
		lineOK("locale", lang)
	} else {
		lineWarn("locale", lang+" — بهتر است *.UTF-8 باشد")
	}

	// Render flags
	if os.Getenv("FAMAN_PLAIN") != "" {
		lineOK("FAMAN_PLAIN", os.Getenv("FAMAN_PLAIN"))
	} else {
		lineOK("FAMAN_PLAIN", "(unset)")
	}

	fmt.Println()
	if warn == 0 {
		if useColor {
			fmt.Println(renderer.SuccessStyle.Render("  همهٔ بررسی‌های اصلی پاس شدند."))
		} else {
			fmt.Println("  all core checks passed")
		}
	} else {
		if useColor {
			fmt.Println(renderer.WarningStyle.Render(fmt.Sprintf("  %d هشدار — docs/install.md و docs/terminal-persian.md", warn)))
		} else {
			fmt.Printf("  %d warning(s)\n", warn)
		}
	}
	_ = filepath.Separator
	return nil
}

func isWritableDir(dir string) bool {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return false
	}
	f, err := os.CreateTemp(dir, ".faman-doctor-*")
	if err != nil {
		return false
	}
	name := f.Name()
	_ = f.Close()
	_ = os.Remove(name)
	return true
}
