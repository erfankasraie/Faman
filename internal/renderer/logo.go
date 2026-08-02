package renderer

import (
	"fmt"
	"os"
	"strings"
)

// Official terminal mark (see assets/logo/terminal-glyph.txt).
// Pages must NOT embed the logo — the CLI always injects it.
const (
	// GlyphOne is the single-cell identity for tight headers.
	GlyphOne = "∧"

	// GlyphBanner is the compact 2-line silhouette (horns + chin).
	GlyphBanner = "   ╭╮  ╭╮\n    ╲▾╱"
)

// PrintBanner writes the logo once (help/version). Respects NO_COLOR / plain.
func PrintBanner(useColor bool) {
	if plainMode() || !useColor {
		fmt.Println(GlyphBanner)
		fmt.Println("  faman — صفحات راهنمای فارسی لینوکس")
		return
	}
	fmt.Println(GlyphStyle.Render(GlyphBanner))
	fmt.Println(DimStyle.Render("  faman — صفحات راهنمای فارسی لینوکس"))
}

// headerTitleLine is "∧  <title>" for every page.
func headerTitleLine(title string) string {
	t := strings.TrimSpace(title)
	if t == "" {
		t = "faman"
	}
	return fmt.Sprintf("%s  %s", GlyphOne, t)
}

// logoDisabled allows tests / tiny CI terminals to skip multi-line art.
func logoBannerEnabled() bool {
	v := strings.ToLower(strings.TrimSpace(os.Getenv("FAMAN_LOGO")))
	if v == "0" || v == "false" || v == "no" || v == "off" {
		return false
	}
	// default: show single-char in header only; full banner on help/version
	return true
}
