package search

import (
	"strings"
	"unicode"

	"github.com/faman-project/faman/internal/parser"
)

// Result is a single search hit.
type Result struct {
	Title    string
	Category string
	Snippet  string
	Score    float64
}

// Options filters search results.
type Options struct {
	Category string // exact category match (case-insensitive); empty = all
}

// Search finds pages matching the query.
func Search(query string) ([]Result, error) {
	return SearchOpts(query, Options{})
}

// SearchOpts finds pages matching the query with optional filters.
func SearchOpts(query string, opt Options) ([]Result, error) {
	pages, err := parser.ListPages()
	if err != nil {
		return nil, err
	}

	catFilter := strings.ToLower(strings.TrimSpace(opt.Category))

	q := normalize(query)
	terms := strings.Fields(q)
	if len(terms) == 0 {
		return nil, nil
	}

	var results []Result
	for _, p := range pages {
		if catFilter != "" && strings.ToLower(p.Category) != catFilter {
			continue
		}
		score, snippet := scorePage(p, terms)
		if score > 0 {
			results = append(results, Result{
				Title:    p.Title,
				Category: p.Category,
				Snippet:  snippet,
				Score:    score,
			})
		}
	}

	for i := 0; i < len(results); i++ {
		for j := i + 1; j < len(results); j++ {
			if results[j].Score > results[i].Score {
				results[i], results[j] = results[j], results[i]
			}
		}
	}
	return results, nil
}

func scorePage(p *parser.Page, terms []string) (float64, string) {
	var score float64
	var snippet string

	titleNorm := normalize(p.Title)
	for _, t := range terms {
		if titleNorm == t {
			score += 10
		} else if strings.Contains(titleNorm, t) {
			score += 5
		}
	}

	for _, a := range p.Aliases {
		an := normalize(a)
		for _, t := range terms {
			if an == t {
				score += 8
			} else if strings.Contains(an, t) {
				score += 3
			}
		}
	}

	for _, k := range p.Keywords {
		kn := normalize(k)
		for _, t := range terms {
			if kn == t {
				score += 4
			} else if strings.Contains(kn, t) {
				score += 2
			}
		}
	}

	contentNorm := normalize(p.Content)
	for _, t := range terms {
		if strings.Contains(contentNorm, t) {
			score += 1
			if snippet == "" {
				snippet = extractSnippet(p.Content, t)
			}
		}
	}

	if p.Category != "" {
		cn := normalize(p.Category)
		for _, t := range terms {
			if strings.Contains(cn, t) {
				score += 1.5
			}
		}
	}

	return score, snippet
}

func normalize(s string) string {
	s = strings.ToLower(s)
	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			b.WriteRune(r)
		}
	}
	return strings.Join(strings.Fields(b.String()), " ")
}

func extractSnippet(content, term string) string {
	lower := strings.ToLower(content)
	idx := strings.Index(lower, strings.ToLower(term))
	if idx < 0 {
		return ""
	}
	start := idx - 40
	if start < 0 {
		start = 0
	}
	end := idx + len(term) + 40
	if end > len(content) {
		end = len(content)
	}
	snip := strings.ReplaceAll(content[start:end], "\n", " ")
	snip = strings.TrimSpace(snip)
	if start > 0 {
		snip = "…" + snip
	}
	if end < len(content) {
		snip = snip + "…"
	}
	return snip
}
