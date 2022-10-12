package _switch

import (
	"github.com/gabe565/gh-profile/internal/profile"
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:     "switch profile",
	Aliases: []string{"activate", "active", "s"},
	Short:   "Switch active profile",
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

	if err := profile.ExistingToDefault(); err != nil {
		return err
	}

	if err := p.Activate(); err != nil {
		return err
	}

	return nil
}
