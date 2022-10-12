package _delete

import (
	"errors"
	"github.com/gabe565/gh-profile/internal/profile"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:     "delete [name]",
	Aliases: []string{"remove", "rm", "d"},
	Short:   "Deletes a profile",
	RunE:    run,
}

func run(cmd *cobra.Command, args []string) (err error) {
	var p profile.Profile
	if len(args) > 0 {
		p = profile.New(args[0])
	} else {
		if p, err = profile.Prompt(); err != nil {
			return err
		}
	}

	if p.IsActive() {
		return errors.New("refusing to delete the active profile. please switch profiles and try again")
	}

	return p.Delete()
}
