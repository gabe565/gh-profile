package util

import (
	"github.com/spf13/cobra"
)

func ShellCompDisable(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return nil, cobra.ShellCompDirectiveNoFileComp
}
