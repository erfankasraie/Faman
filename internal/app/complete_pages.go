package app

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/faman-project/faman/internal/parser"
)

func init() {
	rootCmd.ValidArgsFunction = completePageNames
	searchCmd.ValidArgsFunction = completeSearchQuery
	listCmd.RegisterFlagCompletionFunc("cat", completeCategories)
	searchCmd.RegisterFlagCompletionFunc("cat", completeCategories)
}

func completePageNames(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) > 0 {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	pages, err := parser.ListPages()
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	prefix := strings.ToLower(toComplete)
	var out []string
	for _, p := range pages {
		t := p.Title
		if prefix == "" || strings.HasPrefix(strings.ToLower(t), prefix) {
			desc := p.Category
			if p.Difficulty != "" {
				if desc != "" {
					desc += " · "
				}
				desc += p.Difficulty
			}
			if desc != "" {
				out = append(out, t+"\t"+desc)
			} else {
				out = append(out, t)
			}
		}
	}
	return out, cobra.ShellCompDirectiveNoFileComp
}

func completeSearchQuery(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	// First word: suggest page titles as starting points
	if len(args) == 0 {
		return completePageNames(cmd, args, toComplete)
	}
	return nil, cobra.ShellCompDirectiveNoFileComp
}

func completeCategories(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	pages, err := parser.ListPages()
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}
	seen := map[string]struct{}{}
	prefix := strings.ToLower(toComplete)
	var out []string
	for _, p := range pages {
		c := strings.TrimSpace(p.Category)
		if c == "" {
			continue
		}
		key := strings.ToLower(c)
		if _, ok := seen[key]; ok {
			continue
		}
		if prefix != "" && !strings.HasPrefix(key, prefix) {
			continue
		}
		seen[key] = struct{}{}
		out = append(out, c)
	}
	return out, cobra.ShellCompDirectiveNoFileComp
}
