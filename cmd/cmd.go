package cmd

import (
	"github.com/gabe565/gh-profile/internal/github"
	"github.com/gabe565/gh-profile/internal/profile"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

var Command = &cobra.Command{
	Use:     "profile [name]",
	PreRunE: preRun,
	RunE:    run,
}

func preRun(cmd *cobra.Command, args []string) error {
	if configDir := github.ConfigDir(); strings.HasPrefix(configDir, "$HOME"+string(os.PathSeparator)) {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		configDir = filepath.Join(home, strings.TrimPrefix(configDir, "$HOME"))
		viper.Set(github.ConfigDirKey, configDir)
	}
	return nil
}

func run(cmd *cobra.Command, args []string) (err error) {
	cmd.SilenceUsage = true

	var p profile.Profile
	if len(args) > 0 {
		p = profile.New(args[0])
	} else {
		if p, err = profile.Prompt(); err != nil {
			return err
		}
	}

	if err := profile.ExistingToDefault(); err != nil {
		return err
	}

	if err := p.Activate(); err != nil {
		return err
	}

	return nil
}
