package app

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/faman-project/faman/internal/renderer"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "نمایش نسخه faman",
	Run: func(cmd *cobra.Command, args []string) {
		useColor := renderer.ColorEnabled()
		renderer.PrintBanner(useColor)
		fmt.Println()
		line := fmt.Sprintf("%s  faman  %s", renderer.GlyphOne, version)
		if useColor {
			fmt.Println(renderer.TitleStyle.Render(line))
			fmt.Println(renderer.DimStyle.Render("   Persian Manual Pages"))
		} else {
			fmt.Printf("faman version %s\n", version)
		}
	},
}
