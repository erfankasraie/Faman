package renderer

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-isatty"
)

// Palette — purple/cyan inspired by the project identity
var (
	colorPrimary  = lipgloss.Color("#7C3AED") // violet
	colorAccent   = lipgloss.Color("#06B6D4") // cyan
	colorMuted    = lipgloss.Color("#94A3B8") // slate
	colorSuccess  = lipgloss.Color("#10B981") // emerald
	colorWarning  = lipgloss.Color("#F59E0B") // amber
	colorError    = lipgloss.Color("#EF4444") // red
	colorBeginner = lipgloss.Color("#22C55E")
	colorIntermed = lipgloss.Color("#EAB308")
	colorAdvanced = lipgloss.Color("#EF4444")
	colorBorder   = lipgloss.Color("#4C1D95")
)

// Styles
var (
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorPrimary)

	SubtitleStyle = lipgloss.NewStyle().
			Foreground(colorMuted).
			Italic(true)

	SectionStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorAccent).
			MarginTop(1)

	MetaStyle = lipgloss.NewStyle().
			Foreground(colorMuted)

	AliasStyle = lipgloss.NewStyle().
			Foreground(colorMuted).
			Italic(true)

	BoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(colorPrimary).
			Padding(0, 1).
			MarginBottom(0)

	BadgeBeginner = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#052E16")).
			Background(colorBeginner).
			Padding(0, 1).
			Bold(true)

	BadgeIntermediate = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#422006")).
			Background(colorIntermed).
			Padding(0, 1).
			Bold(true)

	BadgeAdvanced = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#450A0A")).
			Background(colorAdvanced).
			Padding(0, 1).
			Bold(true)

	SearchTitleStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(colorPrimary)

	SearchCatStyle = lipgloss.NewStyle().
			Foreground(colorAccent)

	SearchSnippetStyle = lipgloss.NewStyle().
				Foreground(colorMuted).
				Italic(true)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(colorError).
			Bold(true)

	WarningStyle = lipgloss.NewStyle().
			Foreground(colorWarning).
			Bold(true)

	SuccessStyle = lipgloss.NewStyle().
			Foreground(colorSuccess)

	DimStyle = lipgloss.NewStyle().
			Foreground(colorMuted)

	GlyphStyle = lipgloss.NewStyle().
			Foreground(colorPrimary).
			Bold(true)

	FooterStyle = lipgloss.NewStyle().
			Foreground(colorMuted).
			MarginTop(1)

	// colorBorder reserved for future panels / secondary boxes
	_ = colorBorder
)

// ColorEnabled reports whether ANSI colors should be used.
func ColorEnabled() bool {
	if os.Getenv("NO_COLOR") != "" {
		return false
	}
	if os.Getenv("TERM") == "dumb" {
		return false
	}
	return isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())
}

// DifficultyBadge returns a colored badge for the difficulty level.
func DifficultyBadge(level string) string {
	if !ColorEnabled() {
		return level
	}
	switch level {
	case "beginner", "مبتدی":
		return BadgeBeginner.Render("beginner")
	case "intermediate", "متوسط":
		return BadgeIntermediate.Render("intermediate")
	case "advanced", "پیشرفته":
		return BadgeAdvanced.Render("advanced")
	default:
		return MetaStyle.Render(level)
	}
}

// SectionIcon returns a small icon prefix for known sections.
func SectionIcon(name string) string {
	switch name {
	case "Introduction", "مقدمه":
		return "◈"
	case "Syntax", "نحو", "سینتکس":
		return "⌘"
	case "Options", "گزینه‌ها":
		return "⚙"
	case "Examples", "مثال‌ها":
		return "▸"
	case "Common mistakes", "اشتباهات رایج":
		return "⚠"
	case "Tips", "نکات":
		return "★"
	case "Related commands", "دستورات مرتبط":
		return "→"
	default:
		return "•"
	}
}
