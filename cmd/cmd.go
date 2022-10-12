package cmd

import (
	"github.com/gabe565/gh-profile/cmd/create"
	"github.com/gabe565/gh-profile/cmd/delete"
	"github.com/gabe565/gh-profile/cmd/list"
	"github.com/gabe565/gh-profile/cmd/switch"
	"github.com/gabe565/gh-profile/internal/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

var Command = &cobra.Command{
	Use:               "profile",
	Short:             "Work with multiple GitHub accounts using the gh cli",
	PersistentPreRunE: preRun,
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

	cmd.SilenceUsage = true
	return nil
}

func init() {
	Command.AddCommand(
		create.Command,
		_delete.Command,
		list.Command,
		_switch.Command,
	)
}
