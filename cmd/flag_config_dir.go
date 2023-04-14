package cmd

import (
	"path/filepath"

	ghcliconfig "github.com/cli/go-gh/pkg/config"
	"github.com/gabe565/gh-profile/internal/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func flagConfigDir(cmd *cobra.Command) {
	defaultDir := ghcliconfig.ConfigDir()
	if defaultDir == "" {
			defaultDir = filepath.Join("$HOME", ".config", "gh")
	}
	cmd.PersistentFlags().StringP(github.ConfigDirKey, "c", defaultDir, "GitHub CLI config dir")
	if err := viper.BindPFlag(github.ConfigDirKey, cmd.PersistentFlags().Lookup(github.ConfigDirKey)); err != nil {
		panic(err)
	}
}
