package renderer

import (
	"testing"

	"github.com/faman-project/faman/internal/parser"
)

func TestRenderDoesNotPanic(t *testing.T) {
	page := &parser.Page{
		Title:      "ls",
		Category:   "filesystem",
		Difficulty: "beginner",
		Aliases:    []string{"dir"},
		Content:    "# Introduction\n\nList files.\n\n# Syntax\n\n`ls [options]`\n",
		Sections: map[string]string{
			"Introduction": "List files and directories.",
			"Syntax":       "`ls [options]`",
			"Examples":     "ls -la",
		},
	}
	if err := Render(page); err != nil {
		t.Fatalf("render error: %v", err)
	}
}

func TestDifficultyBadge(t *testing.T) {
	// Should not panic regardless of color support
	_ = DifficultyBadge("beginner")
	_ = DifficultyBadge("intermediate")
	_ = DifficultyBadge("advanced")
	_ = DifficultyBadge("unknown")
}

func TestSectionIcon(t *testing.T) {
	if SectionIcon("Examples") == "" {
		t.Error("expected non-empty icon")
	}
	if SectionIcon("UnknownSection") != "•" {
		t.Error("expected default bullet")
	}
}

func TestMin(t *testing.T) {
	if min(1, 2) != 1 {
		t.Error("min failed")
	}
	if min(5, 3) != 3 {
		t.Error("min failed")
	}
}

func TestColorEnabled(t *testing.T) {
	// Just ensure it doesn't panic
	_ = ColorEnabled()
}
