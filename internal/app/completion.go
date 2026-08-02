package app

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "تولید اسکریپت تکمیل خودکار شل",
	Long: `تولید اسکریپت completion برای شل‌های رایج.

Zsh:
  faman completion zsh > "${fpath[1]}/_faman"
  # یا:
  faman completion zsh > ~/.zsh/completions/_faman
  autoload -Uz compinit && compinit

Bash:
  faman completion bash > /etc/bash_completion.d/faman
  # یا:
  source <(faman completion bash)

Fish:
  faman completion fish > ~/.config/fish/completions/faman.fish
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		switch args[0] {
		case "bash":
			return rootCmd.GenBashCompletionV2(os.Stdout, true)
		case "zsh":
			return rootCmd.GenZshCompletion(os.Stdout)
		case "fish":
			return rootCmd.GenFishCompletion(os.Stdout, true)
		case "powershell":
			return rootCmd.GenPowerShellCompletionWithDesc(os.Stdout)
		default:
			return cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
