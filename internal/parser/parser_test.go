package parser

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSplitFrontMatter(t *testing.T) {
	input := []byte(`---
title: ls
aliases:
- dir
category: filesystem
---
# Introduction

Hello world.
`)
	front, body, err := splitFrontMatter(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(front) == 0 {
		t.Fatal("front matter empty")
	}
	if !contains(string(body), "Introduction") {
		t.Fatalf("body missing content: %s", body)
	}
}

func TestParseFrontMatter(t *testing.T) {
	front := []byte(`title: ls
aliases:
- dir
- list
category: filesystem
difficulty: beginner
keywords:
- files
- folders
`)
	page := &Page{Sections: make(map[string]string)}
	if err := parseFrontMatter(front, page); err != nil {
		t.Fatalf("parse error: %v", err)
	}
	if page.Title != "ls" {
		t.Errorf("title = %q, want ls", page.Title)
	}
	if page.Category != "filesystem" {
		t.Errorf("category = %q", page.Category)
	}
	if page.Difficulty != "beginner" {
		t.Errorf("difficulty = %q", page.Difficulty)
	}
	if len(page.Aliases) != 2 {
		t.Errorf("aliases = %v", page.Aliases)
	}
	if len(page.Keywords) != 2 {
		t.Errorf("keywords = %v", page.Keywords)
	}
}

func TestSplitByHeadings(t *testing.T) {
	content := `# Introduction
This is intro.

# Syntax
ls [options]

# Examples
ls -la
`
	sections := splitByHeadings(content)
	if sections["Introduction"] == "" {
		t.Error("missing Introduction")
	}
	if sections["Syntax"] == "" {
		t.Error("missing Syntax")
	}
	if sections["Examples"] == "" {
		t.Error("missing Examples")
	}
}

func TestLoadPageIntegration(t *testing.T) {
	// Create temporary page
	dir := t.TempDir()
	pageContent := `---
title: testcmd
aliases:
- tc
category: test
difficulty: beginner
keywords:
- test
---
# Introduction

Test page content.
`
	if err := os.WriteFile(filepath.Join(dir, "testcmd.md"), []byte(pageContent), 0644); err != nil {
		t.Fatal(err)
	}

	// Override pagesDir via chdir
	oldWd, _ := os.Getwd()
	// We can't easily override pagesDir, so test parseFile directly
	p, err := parseFile(filepath.Join(dir, "testcmd.md"))
	if err != nil {
		t.Fatal(err)
	}
	if p.Title != "testcmd" {
		t.Errorf("got title %q", p.Title)
	}
	if len(p.Aliases) != 1 || p.Aliases[0] != "tc" {
		t.Errorf("aliases = %v", p.Aliases)
	}
	_ = oldWd
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > 0 && (s[0:len(substr)] == substr || contains(s[1:], substr))))
}
