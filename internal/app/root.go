package app

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/faman-project/faman/internal/parser"
	"github.com/faman-project/faman/internal/renderer"
	"github.com/faman-project/faman/internal/search"
)

var (
	// version is set at link time: -X github.com/faman-project/faman/internal/app.version=...
	version = "0.1.2-pre"
	rootCmd = &cobra.Command{
		Use:   "faman",
		Short: "صفحات راهنمای فارسی لینوکس — Persian Manual Pages",
		Long: `∧  faman — صفحات راهنمای فارسی برای دستورات لینوکس

faman یک ابزار خط فرمان است که دستورات لینوکس را به زبان فارسی
به صورت کامل، کاربردی و مناسب برای مبتدیان و حرفه‌ای‌ها توضیح می‌دهد.

مثال‌ها:
  faman ls
  faman search docker
  faman version`,
		Args: cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return cmd.Help()
			}
			return showPage(args[0])
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}
)

func Execute() error {
	err := rootCmd.Execute()
	if err != nil {
		if renderer.ColorEnabled() {
			fmt.Fprintln(os.Stderr, renderer.ErrorStyle.Render("✗  "+err.Error()))
		} else {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}
	return err
}

func init() {
	rootCmd.AddCommand(searchCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(helpCmd)
	rootCmd.AddCommand(updateCmd)
}

func showPage(name string) error {
	page, err := parser.LoadPage(name)
	if err != nil {
		results, searchErr := search.Search(name)
		if searchErr == nil && len(results) > 0 {
			useColor := renderer.ColorEnabled()
			msg := fmt.Sprintf("صفحه «%s» پیدا نشد.", name)
			if useColor {
				fmt.Fprintln(os.Stderr, renderer.ErrorStyle.Render("✗  "+msg))
				fmt.Fprintln(os.Stderr)
				fmt.Fprintln(os.Stderr, renderer.DimStyle.Render("  آیا منظورتان یکی از این‌ها بود؟"))
			} else {
				fmt.Fprintf(os.Stderr, "%s\n\nآیا منظورتان یکی از این‌ها بود؟\n", msg)
			}
			for _, r := range results {
				if r.Score > 0.3 {
					if useColor {
						fmt.Fprintf(os.Stderr, "    %s  %s\n",
							renderer.SearchTitleStyle.Render("• "+r.Title),
							renderer.SearchCatStyle.Render(r.Category),
						)
					} else {
						fmt.Fprintf(os.Stderr, "  • %s\n", r.Title)
					}
				}
			}
			return fmt.Errorf("صفحه پیدا نشد: %s", name)
		}
		return fmt.Errorf("صفحه پیدا نشد: %s", name)
	}

	return renderer.Render(page)
}
