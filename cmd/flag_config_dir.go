package cmd

import (
	ghcliconfig "github.com/cli/go-gh/v2/pkg/config"
	"github.com/gabe565/gh-profile/internal/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var DefaultConfigDir = ghcliconfig.ConfigDir()

func flagConfigDir(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP(github.ConfigDirKey, "c", DefaultConfigDir, "GitHub CLI config dir")
	if err := viper.BindPFlag(github.ConfigDirKey, cmd.PersistentFlags().Lookup(github.ConfigDirKey)); err != nil {
		panic(err)
	}
	if err := cmd.RegisterFlagCompletionFunc(
		github.ConfigDirKey,
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return nil, cobra.ShellCompDirectiveFilterDirs
		},
	); err != nil {
		panic(err)
	}
}
