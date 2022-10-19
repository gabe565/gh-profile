package cmd

import (
	"github.com/gabe565/gh-profile/internal/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func flagConfigDir(cmd *cobra.Command) {
	defaultDir := os.Getenv("GH_CONFIG_DIR")
	if defaultDir == "" {
		if xdgConfigHome := os.Getenv("XDG_CONFIG_HOME"); xdgConfigHome != "" {
			defaultDir = filepath.Join(xdgConfigHome, "gh")
		} else {
			defaultDir = filepath.Join("$HOME", ".config", "gh")
		}
	}
	cmd.PersistentFlags().StringP(github.ConfigDirKey, "c", defaultDir, "GitHub CLI config dir")
	if err := viper.BindPFlag(github.ConfigDirKey, cmd.PersistentFlags().Lookup(github.ConfigDirKey)); err != nil {
		panic(err)
	}
}
