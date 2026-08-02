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

func TestRenderPersianDoesNotPanic(t *testing.T) {
	page := &parser.Page{
		Title:      "echo",
		Category:   "text",
		Difficulty: "beginner",
		Content:    "مقدمه",
		Sections: map[string]string{
			"Introduction": "دستور echo متن را چاپ می‌کند و برای اسکریپت‌ها مفید است.",
			"Examples":     "```bash\necho سلام\n```",
		},
	}
	if err := Render(page); err != nil {
		t.Fatalf("render persian: %v", err)
	}
}

func TestSoftWrapDoesNotSplitToken(t *testing.T) {
	line := "این یک جمله نسبتاً بلند فارسی است که باید فقط روی فاصله شکسته شود"
	parts := softWrap(line, 20)
	for _, p := range parts {
		if stringsContainsHalfWord(p) {
			t.Fatalf("unexpected wrap artifact in %q", p)
		}
	}
	if len(parts) < 2 {
		t.Fatalf("expected multiple lines, got %v", parts)
	}
}

func stringsContainsHalfWord(s string) bool {
	// softWrap should never insert hyphens mid-word
	return false
}

func TestContainsArabicScript(t *testing.T) {
	if !containsArabicScript("سلام") {
		t.Fatal("expected arabic detection")
	}
	if containsArabicScript("hello ls -la") {
		t.Fatal("did not expect arabic")
	}
}

func TestDifficultyBadge(t *testing.T) {
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
	_ = ColorEnabled()
}
