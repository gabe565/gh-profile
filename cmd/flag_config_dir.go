package cmd

import (
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func init() {
	defaultDir := os.Getenv("GH_CONFIG_DIR")
	if defaultDir == "" {
		defaultDir = filepath.Join("$HOME", ".config", "gh")
	}
	Command.PersistentFlags().StringP("config-dir", "c", defaultDir, "GitHub CLI config dir")
	if err := viper.BindPFlag("config-dir", Command.PersistentFlags().Lookup("config-dir")); err != nil {
		panic(err)
	}
}
