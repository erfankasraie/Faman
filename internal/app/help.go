package app

import (
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help [command]",
	Short: "راهنمای استفاده از faman",
	Long:  `نمایش راهنمای کامل faman یا یک زیر‌دستور خاص.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			_ = rootCmd.Help()
			return
		}
		// Find and show help for subcommand
		c, _, err := rootCmd.Find(args)
		if err != nil || c == nil {
			_ = rootCmd.Help()
			return
		}
		_ = c.Help()
	},
}
