package create

import (
	"github.com/gabe565/gh-profile/internal/profile"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	return &cobra.Command{
		Use:     "create [name]",
		Aliases: []string{"c", "new", "add"},
		Short:   "Creates a new profile",
		RunE:    run,
	}
}

func run(cmd *cobra.Command, args []string) (err error) {
	var p profile.Profile
	if len(args) > 0 {
		p = profile.New(args[0])
	} else {
		if p, err = profile.PromptNew(); err != nil {
			return err
		}
	}

	if err := profile.ExistingToDefault(); err != nil {
		return err
	}

	if err := p.Create(); err != nil {
		return err
	}

	return p.Activate()
}
