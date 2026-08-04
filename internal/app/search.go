package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/faman-project/faman/internal/renderer"
	"github.com/faman-project/faman/internal/search"
)

var searchCategory string

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "جستجو در تمام صفحات راهنما",
	Long: `جستجو در عنوان، کلیدواژه‌ها، نام‌های مستعار و محتوای صفحات راهنما.

  faman search docker
  faman search port --cat network
  faman search file --cat filesystem`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		query := strings.Join(args, " ")
		results, err := search.SearchOpts(query, search.Options{Category: searchCategory})
		if err != nil {
			return err
		}

		useColor := renderer.ColorEnabled()

		if len(results) == 0 {
			msg := fmt.Sprintf("هیچ نتیجه‌ای برای «%s» پیدا نشد.", query)
			if searchCategory != "" {
				msg = fmt.Sprintf("هیچ نتیجه‌ای برای «%s» در دسته «%s» پیدا نشد.", query, searchCategory)
			}
			if useColor {
				fmt.Fprintln(os.Stderr, renderer.ErrorStyle.Render("✗  "+msg))
			} else {
				fmt.Fprintln(os.Stderr, msg)
			}
			return nil
		}

		header := fmt.Sprintf("∧  نتایج جستجو برای «%s»  ·  %d مورد", query, len(results))
		if searchCategory != "" {
			header = fmt.Sprintf("∧  جستجو «%s» در [%s]  ·  %d مورد", query, searchCategory, len(results))
		}
		if useColor {
			fmt.Println(renderer.SearchTitleStyle.Render(header))
			fmt.Println(renderer.DimStyle.Render(strings.Repeat("─", 40)))
		} else {
			fmt.Println(header)
			fmt.Println(strings.Repeat("-", 40))
		}
		fmt.Println()

		limit := 15
		for i, r := range results {
			if i >= limit {
				more := fmt.Sprintf("… و %d نتیجه دیگر", len(results)-limit)
				if useColor {
					fmt.Println(renderer.DimStyle.Render(more))
				} else {
					fmt.Println(more)
				}
				break
			}

			title := r.Title
			cat := r.Category
			if useColor {
				fmt.Printf("  %s  %s\n",
					renderer.SearchTitleStyle.Render(fmt.Sprintf("%-12s", title)),
					renderer.SearchCatStyle.Render(cat),
				)
			} else {
				fmt.Printf("  %-12s  %s\n", title, cat)
			}

			if r.Snippet != "" {
				snip := truncate(r.Snippet, 70)
				if useColor {
					fmt.Println(renderer.SearchSnippetStyle.Render("               " + snip))
				} else {
					fmt.Printf("               %s\n", snip)
				}
			}
		}

		fmt.Println()
		hint := "برای مشاهده:  faman <command>"
		if useColor {
			fmt.Println(renderer.DimStyle.Render(hint))
		} else {
			fmt.Println(hint)
		}
		return nil
	},
}

func init() {
	searchCmd.Flags().StringVar(&searchCategory, "cat", "", "محدود کردن جستجو به یک دسته (مثل network)")
}

func truncate(s string, n int) string {
	s = strings.Join(strings.Fields(s), " ")
	runes := []rune(s)
	if len(runes) <= n {
		return s
	}
	return string(runes[:n-1]) + "…"
}
