package parser

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

// Page represents a single manual page.
type Page struct {
	Title      string
	Aliases    []string
	Category   string
	Keywords   []string
	Difficulty string
	Content    string // raw markdown body
	Sections   map[string]string
	FilePath   string
}

// pagesDir returns the directory containing markdown pages.
// Search order:
//  1. FAMAN_PAGES env var
//  2. Next to the executable
//  3. Standard system share directories
//  4. Current working directory (development)
func pagesDir() (string, error) {
	// 1. Explicit environment variable
	if env := os.Getenv("FAMAN_PAGES"); env != "" {
		if st, err := os.Stat(env); err == nil && st.IsDir() {
			return env, nil
		}
	}

	var candidates []string

	// 2. Relative to executable
	if exe, err := os.Executable(); err == nil {
		exeDir := filepath.Dir(exe)
		candidates = append(candidates,
			filepath.Join(exeDir, "pages", "fa"),
			filepath.Join(exeDir, "..", "pages", "fa"),
			filepath.Join(exeDir, "..", "share", "faman", "pages", "fa"),
		)
	}

	// 3. Standard FHS locations
	candidates = append(candidates,
		"/usr/local/share/faman/pages/fa",
		"/usr/share/faman/pages/fa",
		"/opt/faman/pages/fa",
	)

	// 4. Development / current directory
	if cwd, err := os.Getwd(); err == nil {
		candidates = append(candidates,
			filepath.Join(cwd, "pages", "fa"),
			filepath.Join(cwd, "..", "pages", "fa"),
		)
	}

	for _, dir := range candidates {
		if st, err := os.Stat(dir); err == nil && st.IsDir() {
			return dir, nil
		}
	}
	return "", fmt.Errorf("pages directory not found (set FAMAN_PAGES or install pages to /usr/local/share/faman/pages/fa)")
}

// LoadPage loads a page by name or alias.
func LoadPage(name string) (*Page, error) {
	dir, err := pagesDir()
	if err != nil {
		return nil, err
	}

	name = strings.ToLower(strings.TrimSpace(name))

	// Direct file match
	path := filepath.Join(dir, name+".md")
	if _, err := os.Stat(path); err == nil {
		return parseFile(path)
	}

	// Search by alias
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		p, err := parseFile(filepath.Join(dir, e.Name()))
		if err != nil {
			continue
		}
		if strings.EqualFold(p.Title, name) {
			return p, nil
		}
		for _, a := range p.Aliases {
			if strings.EqualFold(a, name) {
				return p, nil
			}
		}
	}
	return nil, fmt.Errorf("page not found: %s", name)
}

// ListPages returns all available pages.
func ListPages() ([]*Page, error) {
	dir, err := pagesDir()
	if err != nil {
		return nil, err
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var pages []*Page
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		p, err := parseFile(filepath.Join(dir, e.Name()))
		if err != nil {
			continue
		}
		pages = append(pages, p)
	}
	return pages, nil
}

func parseFile(path string) (*Page, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	front, body, err := splitFrontMatter(data)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}

	page := &Page{
		Content:  string(body),
		Sections: make(map[string]string),
		FilePath: path,
	}

	if err := parseFrontMatter(front, page); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}

	// Extract sections
	extractSections(body, page)

	return page, nil
}

func splitFrontMatter(data []byte) (front, body []byte, err error) {
	const delim = "---"
	if !bytes.HasPrefix(data, []byte(delim)) {
		return nil, data, nil
	}
	rest := data[len(delim):]
	if len(rest) > 0 && rest[0] == '\n' {
		rest = rest[1:]
	} else if len(rest) > 1 && rest[0] == '\r' && rest[1] == '\n' {
		rest = rest[2:]
	}

	idx := bytes.Index(rest, []byte("\n"+delim))
	if idx < 0 {
		return nil, nil, fmt.Errorf("unclosed front matter")
	}
	front = bytes.TrimSpace(rest[:idx])
	body = bytes.TrimSpace(rest[idx+len("\n"+delim):])
	if len(body) > 0 && body[0] == '\n' {
		body = body[1:]
	}
	return front, body, nil
}

func parseFrontMatter(data []byte, page *Page) error {
	lines := strings.Split(string(data), "\n")
	var currentKey string
	var listMode bool

	for _, line := range lines {
		line = strings.TrimRight(line, "\r")
		trimmed := strings.TrimSpace(line)

		if trimmed == "" {
			continue
		}

		if strings.HasPrefix(trimmed, "- ") {
			if !listMode || currentKey == "" {
				continue
			}
			val := strings.TrimSpace(trimmed[2:])
			switch currentKey {
			case "aliases":
				page.Aliases = append(page.Aliases, val)
			case "keywords":
				page.Keywords = append(page.Keywords, val)
			}
			continue
		}

		parts := strings.SplitN(trimmed, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		currentKey = key
		listMode = val == ""

		switch key {
		case "title":
			page.Title = val
		case "category":
			page.Category = val
		case "difficulty":
			page.Difficulty = val
		case "aliases", "keywords":
			// list follows
		}
	}
	return nil
}

func extractSections(body []byte, page *Page) {
	md := goldmark.New()
	reader := text.NewReader(body)
	doc := md.Parser().Parse(reader)

	var currentSection string
	var buf strings.Builder

	flush := func() {
		if currentSection != "" {
			page.Sections[currentSection] = strings.TrimSpace(buf.String())
			buf.Reset()
		}
	}

	// Walk always returns nil error with this visitor; still check for errcheck.
	_ = ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}
		if h, ok := n.(*ast.Heading); ok {
			if h.Level == 1 {
				flush()
				var title strings.Builder
				for c := h.FirstChild(); c != nil; c = c.NextSibling() {
					if t, ok := c.(*ast.Text); ok {
						title.Write(t.Segment.Value(body))
					}
				}
				currentSection = strings.TrimSpace(title.String())
				return ast.WalkSkipChildren, nil
			}
		}
		return ast.WalkContinue, nil
	})
	flush()

	if len(page.Sections) == 0 || allEmpty(page.Sections) {
		page.Sections = splitByHeadings(string(body))
	}
}

func allEmpty(m map[string]string) bool {
	for _, v := range m {
		if strings.TrimSpace(v) != "" {
			return false
		}
	}
	return true
}

func splitByHeadings(content string) map[string]string {
	sections := make(map[string]string)
	lines := strings.Split(content, "\n")
	var current string
	var buf strings.Builder

	flush := func() {
		if current != "" {
			sections[current] = strings.TrimSpace(buf.String())
			buf.Reset()
		}
	}

	for _, line := range lines {
		if strings.HasPrefix(line, "# ") {
			flush()
			current = strings.TrimSpace(line[2:])
			continue
		}
		if current != "" {
			buf.WriteString(line)
			buf.WriteByte('\n')
		}
	}
	flush()
	return sections
}
