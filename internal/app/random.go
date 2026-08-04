package app

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"

	"github.com/faman-project/faman/internal/parser"
	"github.com/faman-project/faman/internal/renderer"
)

var randomOpen bool

var randomCmd = &cobra.Command{
	Use:     "random",
	Aliases: []string{"rand"},
	Short:   "یک صفحهٔ تصادفی برای یادگیری",
	Long: `نام (و در صورت --open متن) یک صفحهٔ تصادفی را نشان می‌دهد.

  faman random
  faman random --open`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runRandom()
	},
}

func init() {
	randomCmd.Flags().BoolVar(&randomOpen, "open", false, "نمایش کامل صفحه")
}

func runRandom() error {
	pages, err := parser.ListPages()
	if err != nil {
		return err
	}
	if len(pages) == 0 {
		return fmt.Errorf("هیچ صفحه‌ای یافت نشد")
	}
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	p := pages[rng.Intn(len(pages))]

	if randomOpen {
		return renderer.Render(p)
	}
	useColor := renderer.ColorEnabled()
	if useColor {
		fmt.Println(renderer.TitleStyle.Render("∧  صفحهٔ تصادفی"))
		fmt.Printf("  %s  %s\n",
			renderer.SearchTitleStyle.Render(p.Title),
			renderer.SearchCatStyle.Render(p.Category),
		)
		fmt.Println(renderer.DimStyle.Render("  faman " + p.Title + "   یا   faman random --open"))
	} else {
		fmt.Printf("تصادفی: %s (%s)\n", p.Title, p.Category)
		fmt.Printf("  faman %s\n", p.Title)
	}
	return nil
}
