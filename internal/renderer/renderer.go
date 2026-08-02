package renderer

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-runewidth"
	"golang.org/x/term"

	"github.com/faman-project/faman/internal/parser"
)

// Render displays a page beautifully in the terminal.
// Persian/Arabic text is sensitive to mid-word wrapping (breaks glyph joining).
// We soft-wrap only on spaces and offer FAMAN_PLAIN / FAMAN_WRAP controls.
func Render(page *parser.Page) error {
	warnIfBadLocale()

	width := termWidth()
	useColor := ColorEnabled()
	plain := plainMode()

	renderHeader(page, width, useColor && !plain)

	order := []string{
		"Introduction", "مقدمه",
		"Syntax", "نحو", "سینتکس",
		"Options", "گزینه‌ها",
		"Examples", "مثال‌ها",
		"Common mistakes", "اشتباهات رایج",
		"Tips", "نکات",
		"Related commands", "دستورات مرتبط",
	}

	rendered := make(map[string]bool)

	var mdRenderer *glamour.TermRenderer
	if useColor && !plain && wrapEnabled() {
		// Keep wrap width generous; Persian words must not be split mid-token.
		wr := width - 2
		if wr < 40 {
			wr = 40
		}
		if wr > 100 {
			wr = 100
		}
		r, err := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
			glamour.WithWordWrap(wr),
			glamour.WithPreservedNewLines(),
		)
		if err == nil {
			mdRenderer = r
		}
	}

	printSection := func(key, content string) {
		icon := SectionIcon(key)
		label := fmt.Sprintf("%s  %s", icon, key)
		if useColor && !plain {
			fmt.Println(SectionStyle.Render(label))
		} else {
			fmt.Printf("\n## %s\n", key)
		}

		body := strings.TrimSpace(content)
		if plain || !wrapEnabled() {
			printPlainBody(body)
			return
		}
		if mdRenderer != nil && !containsArabicScript(body) {
			// Latin-heavy sections (code-only) can use glamour safely.
			out, err := mdRenderer.Render(body)
			if err == nil {
				fmt.Print(out)
				return
			}
		}
		// Persian-aware path: wrap only at spaces, preserve code fences.
		printPersianSafe(body, width-2)
	}

	for _, key := range order {
		content, ok := page.Sections[key]
		if !ok || strings.TrimSpace(content) == "" {
			continue
		}
		if rendered[key] {
			continue
		}
		rendered[key] = true
		printSection(key, content)
	}

	for key, content := range page.Sections {
		if rendered[key] || strings.TrimSpace(content) == "" {
			continue
		}
		printSection(key, content)
	}

	renderFooter(page, useColor && !plain)
	return nil
}

func plainMode() bool {
	v := strings.ToLower(strings.TrimSpace(os.Getenv("FAMAN_PLAIN")))
	return v == "1" || v == "true" || v == "yes"
}

func wrapEnabled() bool {
	v := strings.ToLower(strings.TrimSpace(os.Getenv("FAMAN_WRAP")))
	if v == "0" || v == "false" || v == "no" || v == "off" {
		return false
	}
	return true
}

func warnIfBadLocale() {
	lang := os.Getenv("LANG")
	lc := os.Getenv("LC_ALL")
	lcctype := os.Getenv("LC_CTYPE")
	combined := strings.ToLower(lang + " " + lc + " " + lcctype)
	if combined == "" {
		return
	}
	if strings.Contains(combined, "utf-8") || strings.Contains(combined, "utf8") {
		return
	}
	// Non-UTF-8 locales often show Persian as mojibake (توهم حروف).
	fmt.Fprintln(os.Stderr, "faman: هشدار — locale شما UTF-8 نیست؛ ممکن است حروف فارسی خراب دیده شوند.")
	fmt.Fprintln(os.Stderr, "  پیشنهاد: export LANG=en_US.UTF-8   یا   export LC_ALL=fa_IR.UTF-8")
	fmt.Fprintln(os.Stderr, "  راهنما: docs/terminal-persian.md  |  حالت ساده: FAMAN_PLAIN=1 faman <cmd>")
}

func containsArabicScript(s string) bool {
	for _, r := range s {
		if unicode.In(r, unicode.Arabic, unicode.Inherited) {
			return true
		}
		// Persian additions in Arabic block are covered; also check presentation forms range.
		if r >= 0x0600 && r <= 0x06FF {
			return true
		}
		if r >= 0xFB50 && r <= 0xFDFF {
			return true
		}
		if r >= 0xFE70 && r <= 0xFEFF {
			return true
		}
	}
	return false
}

func printPlainBody(body string) {
	for _, line := range strings.Split(body, "\n") {
		fmt.Println(line)
	}
	fmt.Println()
}

// printPersianSafe prints markdown-ish body without splitting Persian words.
func printPersianSafe(body string, width int) {
	if width < 20 {
		width = 20
	}
	lines := strings.Split(body, "\n")
	inFence := false
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "```") {
			inFence = !inFence
			fmt.Println(line)
			continue
		}
		if inFence || trimmed == "" {
			fmt.Println(line)
			continue
		}
		// Tables and headings: print as-is (avoid breaking | columns).
		if strings.HasPrefix(trimmed, "|") || strings.HasPrefix(trimmed, "#") {
			fmt.Println(line)
			continue
		}
		for _, wrapped := range softWrap(line, width) {
			fmt.Println(wrapped)
		}
	}
	fmt.Println()
}

// softWrap breaks only on whitespace; never mid-token (protects Arabic joining).
func softWrap(s string, width int) []string {
	if runewidth.StringWidth(s) <= width {
		return []string{s}
	}
	words := strings.Fields(s)
	if len(words) == 0 {
		return []string{s}
	}
	var out []string
	var cur strings.Builder
	curWidth := 0
	for _, w := range words {
		ww := runewidth.StringWidth(w)
		if curWidth == 0 {
			cur.WriteString(w)
			curWidth = ww
			continue
		}
		if curWidth+1+ww <= width {
			cur.WriteByte(' ')
			cur.WriteString(w)
			curWidth += 1 + ww
			continue
		}
		out = append(out, cur.String())
		cur.Reset()
		cur.WriteString(w)
		curWidth = ww
	}
	if cur.Len() > 0 {
		out = append(out, cur.String())
	}
	return out
}

func renderHeader(page *parser.Page, width int, useColor bool) {
	glyph := "∧"
	titleLine := fmt.Sprintf("%s  %s", glyph, page.Title)

	if useColor {
		// Avoid fixed Width() — miscalculated for some Unicode fonts and clips text.
		box := BoxStyle.Render(TitleStyle.Render(titleLine))
		fmt.Println(box)
	} else {
		w := min(width, 48)
		fmt.Println(strings.Repeat("─", w))
		fmt.Printf("  %s\n", titleLine)
		fmt.Println(strings.Repeat("─", w))
	}

	badge := DifficultyBadge(page.Difficulty)
	meta := fmt.Sprintf("%s  ·  %s", page.Category, badge)
	if useColor {
		fmt.Println(MetaStyle.Render(meta))
	} else {
		fmt.Println(meta)
	}

	if len(page.Aliases) > 0 {
		aliases := "aliases: " + strings.Join(page.Aliases, " · ")
		if useColor {
			fmt.Println(AliasStyle.Render(aliases))
		} else {
			fmt.Println(aliases)
		}
	}

	divLen := min(width-2, 48)
	if divLen < 8 {
		divLen = 8
	}
	if useColor {
		div := lipgloss.NewStyle().Foreground(colorBorder).Render(strings.Repeat("─", divLen))
		fmt.Println(div)
	} else {
		fmt.Println(strings.Repeat("-", divLen))
	}
	fmt.Println()
}

func renderFooter(page *parser.Page, useColor bool) {
	hint := fmt.Sprintf("faman search %s  ·  faman help", page.Title)
	if useColor {
		fmt.Println(FooterStyle.Render("─────────────────────────────────────"))
		fmt.Println(FooterStyle.Render(hint))
	} else {
		fmt.Println("-------------------------------------")
		fmt.Println(hint)
	}
}

func termWidth() int {
	w, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || w < 40 {
		return 80
	}
	return w
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
