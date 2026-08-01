package search

import (
	"testing"

	"github.com/faman-project/faman/internal/parser"
)

func TestNormalize(t *testing.T) {
	got := normalize("  Hello, World! 123 ")
	want := "hello world 123"
	if got != want {
		t.Errorf("normalize = %q, want %q", got, want)
	}
}

func TestScorePage(t *testing.T) {
	p := &parser.Page{
		Title:    "ls",
		Aliases:  []string{"dir"},
		Keywords: []string{"files", "list"},
		Category: "filesystem",
		Content:  "List directory contents. Use ls -la for details.",
	}
	score, snip := scorePage(p, []string{"ls"})
	if score < 5 {
		t.Errorf("expected high score for exact title, got %v", score)
	}
	_ = snip

	score2, _ := scorePage(p, []string{"dir"})
	if score2 < 5 {
		t.Errorf("expected high score for alias, got %v", score2)
	}

	score3, _ := scorePage(p, []string{"nonexistentxyz"})
	if score3 > 0 {
		t.Errorf("expected zero score, got %v", score3)
	}
}

func TestExtractSnippet(t *testing.T) {
	content := "This is a long piece of text that contains the word docker somewhere in the middle of the sentence for testing."
	snip := extractSnippet(content, "docker")
	if snip == "" {
		t.Error("expected non-empty snippet")
	}
	if !contains(snip, "docker") {
		t.Errorf("snippet should contain term: %s", snip)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || indexOf(s, substr) >= 0)
}

func indexOf(s, substr string) int {
	for i := 0; i+len(substr) <= len(s); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
