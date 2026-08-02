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
// Logo is injected here — never required inside markdown translations.
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
			out, err := mdRenderer.Render(body)
			if err == nil {
				fmt.Print(out)
				return
			}
		}
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

func containsArabicScript(s string) bool {
	for _, r := range s {
		if unicode.In(r, unicode.Arabic, unicode.Inherited) {
			return true
		}
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
	titleLine := headerTitleLine(page.Title)

	if useColor {
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
	hint := fmt.Sprintf("%s  faman search %s  ·  faman help", GlyphOne, page.Title)
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
