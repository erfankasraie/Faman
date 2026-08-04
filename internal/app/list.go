package app

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/spf13/cobra"

	"github.com/faman-project/faman/internal/parser"
	"github.com/faman-project/faman/internal/renderer"
)

var (
	listCategory   string
	listDifficulty string
	listNamesOnly  bool
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls-pages", "pages"},
	Short:   "فهرست صفحات راهنما",
	Long: `فهرست همهٔ صفحات موجود با فیلتر اختیاری دسته و سطح.

  faman list
  faman list --cat network
  faman list --diff beginner
  faman list --names`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runList()
	},
}

func init() {
	listCmd.Flags().StringVar(&listCategory, "cat", "", "فیلتر دسته (category)")
	listCmd.Flags().StringVar(&listDifficulty, "diff", "", "فیلتر سطح: beginner|intermediate|advanced")
	listCmd.Flags().BoolVar(&listNamesOnly, "names", false, "فقط نام صفحات")
}

func runList() error {
	pages, err := parser.ListPages()
	if err != nil {
		return err
	}
	catFilter := strings.ToLower(strings.TrimSpace(listCategory))
	diffFilter := strings.ToLower(strings.TrimSpace(listDifficulty))

	var out []*parser.Page
	for _, p := range pages {
		if catFilter != "" && strings.ToLower(p.Category) != catFilter {
			continue
		}
		if diffFilter != "" && strings.ToLower(p.Difficulty) != diffFilter {
			continue
		}
		out = append(out, p)
	}
	sort.Slice(out, func(i, j int) bool {
		return strings.ToLower(out[i].Title) < strings.ToLower(out[j].Title)
	})

	useColor := renderer.ColorEnabled()
	if useColor {
		fmt.Println(renderer.TitleStyle.Render(fmt.Sprintf("∧  صفحات faman  ·  %d مورد", len(out))))
		fmt.Println(renderer.DimStyle.Render(strings.Repeat("─", 40)))
	} else {
		fmt.Printf("صفحات faman (%d)\n", len(out))
	}

	if len(out) == 0 {
		fmt.Fprintln(os.Stderr, "موردی با این فیلتر پیدا نشد.")
		return nil
	}

	for _, p := range out {
		if listNamesOnly {
			fmt.Println(p.Title)
			continue
		}
		cat := p.Category
		if cat == "" {
			cat = "-"
		}
		diff := p.Difficulty
		if diff == "" {
			diff = "-"
		}
		if useColor {
			fmt.Printf("  %s  %s  %s\n",
				renderer.SearchTitleStyle.Render(fmt.Sprintf("%-18s", p.Title)),
				renderer.SearchCatStyle.Render(fmt.Sprintf("%-14s", cat)),
				renderer.DimStyle.Render(diff),
			)
		} else {
			fmt.Printf("  %-18s  %-14s  %s\n", p.Title, cat, diff)
		}
	}
	return nil
}
