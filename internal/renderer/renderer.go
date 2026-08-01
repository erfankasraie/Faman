package renderer

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"

	"github.com/faman-project/faman/internal/parser"
)

// Render displays a page beautifully in the terminal.
func Render(page *parser.Page) error {
	width := termWidth()
	useColor := ColorEnabled()

	// ── Header ──────────────────────────────────────────────
	renderHeader(page, width, useColor)

	// ── Sections ────────────────────────────────────────────
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
	if useColor {
		r, err := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
			glamour.WithWordWrap(width-4),
		)
		if err == nil {
			mdRenderer = r
		}
	}

	printSection := func(key, content string) {
		icon := SectionIcon(key)
		label := fmt.Sprintf("%s  %s", icon, key)
		if useColor {
			fmt.Println(SectionStyle.Render(label))
		} else {
			fmt.Printf("\n## %s\n", key)
		}

		body := strings.TrimSpace(content)
		if mdRenderer != nil {
			out, err := mdRenderer.Render(body)
			if err == nil {
				fmt.Print(out)
				return
			}
		}
		// Plain fallback
		for _, line := range strings.Split(body, "\n") {
			fmt.Println("  " + line)
		}
		fmt.Println()
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

	// Remaining sections
	for key, content := range page.Sections {
		if rendered[key] || strings.TrimSpace(content) == "" {
			continue
		}
		printSection(key, content)
	}

	// ── Footer ──────────────────────────────────────────────
	renderFooter(page, useColor)
	return nil
}

func renderHeader(page *parser.Page, width int, useColor bool) {
	// Glyph + title line
	glyph := "∧"
	titleLine := fmt.Sprintf("%s  %s", glyph, page.Title)

	if useColor {
		box := BoxStyle.Width(min(width-2, 48)).Render(TitleStyle.Render(titleLine))
		fmt.Println(box)
	} else {
		fmt.Println(strings.Repeat("─", min(width, 48)))
		fmt.Printf("  %s\n", titleLine)
		fmt.Println(strings.Repeat("─", min(width, 48)))
	}

	// Meta line: category • difficulty badge
	badge := DifficultyBadge(page.Difficulty)
	meta := fmt.Sprintf("%s  ·  %s", page.Category, badge)
	if useColor {
		fmt.Println(MetaStyle.Render(meta))
	} else {
		fmt.Println(meta)
	}

	// Aliases
	if len(page.Aliases) > 0 {
		aliases := "aliases: " + strings.Join(page.Aliases, " · ")
		if useColor {
			fmt.Println(AliasStyle.Render(aliases))
		} else {
			fmt.Println(aliases)
		}
	}

	// Divider
	if useColor {
		div := lipgloss.NewStyle().Foreground(colorBorder).Render(strings.Repeat("─", min(width-2, 48)))
		fmt.Println(div)
	} else {
		fmt.Println(strings.Repeat("-", min(width, 48)))
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

func renderPlain(page *parser.Page) error {
	fmt.Printf("# %s\n\n", page.Title)
	fmt.Println(page.Content)
	return nil
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
