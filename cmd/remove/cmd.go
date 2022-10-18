package remove

import (
	"errors"
	"github.com/gabe565/gh-profile/internal/profile"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	return &cobra.Command{
		Use:     "remove [NAME]",
		Aliases: []string{"delete", "rm", "d"},
		Short:   "Deletes a profile",
		RunE:    run,
	}
}

func run(cmd *cobra.Command, args []string) (err error) {
	var p profile.Profile
	if len(args) > 0 {
		p = profile.New(args[0])
	} else {
		if p, err = profile.Select("Choose a profile to remove"); err != nil {
			return err
		}
	}

	if p.IsActive() {
		return errors.New("refusing to remove the active profile. please switch profiles and try again")
	}

	return p.Remove()
}
