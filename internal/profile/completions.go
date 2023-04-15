package profile

import (
	"github.com/spf13/cobra"
)

func ShellCompName(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	if len(args) > 0 {
		return []string{}, cobra.ShellCompDirectiveNoFileComp
	}

	profiles, _ := List()
	names := make([]string, 0, len(profiles))
	for _, p := range profiles {
		names = append(names, p.Name)
	}
	return names, cobra.ShellCompDirectiveNoFileComp
}
