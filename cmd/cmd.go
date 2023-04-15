package cmd

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/gabe565/gh-profile/cmd/create"
	"github.com/gabe565/gh-profile/cmd/list"
	"github.com/gabe565/gh-profile/cmd/remove"
	"github.com/gabe565/gh-profile/cmd/rename"
	"github.com/gabe565/gh-profile/cmd/show"
	"github.com/gabe565/gh-profile/cmd/switch"
	"github.com/gabe565/gh-profile/internal/github"
	"github.com/gabe565/gh-profile/internal/profile"
	"github.com/spf13/cobra"
)

func New(version, commit string) *cobra.Command {
	cmd := &cobra.Command{
		Use:               "profile",
		Short:             "Work with multiple GitHub accounts using the gh cli",
		PersistentPreRunE: preRun,
		DisableAutoGenTag: true,
		SilenceErrors:     true,
		Version:           buildVersion(version, commit),
	}
	cmd.AddCommand(
		create.New(),
		remove.New(),
		list.New(),
		_switch.New(),
		rename.New(),
		show.New(),
	)
	flagConfigDir(cmd)
	return cmd
}

func preRun(cmd *cobra.Command, args []string) error {
	configDir := github.ConfigDir()
	if strings.HasPrefix(configDir, "$HOME"+string(os.PathSeparator)) {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		configDir = filepath.Join(home, strings.TrimPrefix(configDir, "$HOME"))
		github.SetConfigDir(configDir)
	}

	if dir := filepath.Dir(configDir); filepath.Base(dir) == profile.ConfigDirName {
		github.SetRootConfigDir(filepath.Dir(dir))
	} else {
		github.SetRootConfigDir(configDir)
	}

	cmd.SilenceUsage = true
	return nil
}

func buildVersion(version, commit string) string {
	if commit != "" {
		version += " (" + commit + ")"
	}
	return version
}
