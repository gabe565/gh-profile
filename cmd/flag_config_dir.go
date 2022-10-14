package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func flagConfigDir(cmd *cobra.Command) {
	defaultDir := os.Getenv("GH_CONFIG_DIR")
	if defaultDir == "" {
		defaultDir = filepath.Join("$HOME", ".config", "gh")
	}
	cmd.PersistentFlags().StringP("config-dir", "c", defaultDir, "GitHub CLI config dir")
	if err := viper.BindPFlag("config-dir", cmd.PersistentFlags().Lookup("config-dir")); err != nil {
		panic(err)
	}
}
