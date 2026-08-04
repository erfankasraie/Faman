package app

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"

	"github.com/faman-project/faman/internal/parser"
	"github.com/faman-project/faman/internal/renderer"
)

var categoriesCmd = &cobra.Command{
	Use:     "categories",
	Aliases: []string{"cats"},
	Short:   "دسته‌بندی صفحات و تعداد هر کدام",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runCategories()
	},
}

func runCategories() error {
	pages, err := parser.ListPages()
	if err != nil {
		return err
	}
	counts := map[string]int{}
	for _, p := range pages {
		c := strings.TrimSpace(p.Category)
		if c == "" {
			c = "(بدون دسته)"
		}
		counts[c]++
	}
	names := make([]string, 0, len(counts))
	for k := range counts {
		names = append(names, k)
	}
	sort.Strings(names)

	useColor := renderer.ColorEnabled()
	if useColor {
		fmt.Println(renderer.TitleStyle.Render(fmt.Sprintf("∧  دسته‌ها  ·  %d صفحه در %d دسته", len(pages), len(names))))
		fmt.Println(renderer.DimStyle.Render(strings.Repeat("─", 40)))
	} else {
		fmt.Printf("دسته‌ها (%d صفحه)\n", len(pages))
	}
	for _, n := range names {
		if useColor {
			fmt.Printf("  %s  %s\n",
				renderer.SearchTitleStyle.Render(fmt.Sprintf("%-20s", n)),
				renderer.MetaStyle.Render(fmt.Sprintf("%d", counts[n])),
			)
		} else {
			fmt.Printf("  %-20s  %d\n", n, counts[n])
		}
	}
	return nil
}
