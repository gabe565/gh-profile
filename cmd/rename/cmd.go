package rename

import (
	"github.com/gabe565/gh-profile/internal/profile"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:     "rename [name] [new-name]",
	Aliases: []string{"mv"},
	Short:   "Renames a profile",
	RunE:    run,
}

func run(cmd *cobra.Command, args []string) (err error) {
	var oldProfile profile.Profile
	if len(args) > 0 {
		oldProfile = profile.New(args[0])
	} else {
		if oldProfile, err = profile.Select("Choose a profile to rename"); err != nil {
			return err
		}
	}

	var newProfile profile.Profile
	if len(args) > 1 {
		newProfile = profile.New(args[1])
	} else {
		if newProfile, err = profile.PromptNew(); err != nil {
			return err
		}
	}

	return oldProfile.Rename(newProfile.Name)
}
