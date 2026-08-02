package app

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/faman-project/faman/internal/renderer"
)

var helpCmd = &cobra.Command{
	Use:   "help [command]",
	Short: "راهنمای faman",
	Run: func(cmd *cobra.Command, args []string) {
		useColor := renderer.ColorEnabled()
		renderer.PrintBanner(useColor)
		fmt.Println()
		_ = rootCmd.Help()
	},
}
