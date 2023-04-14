package _switch

import (
	"errors"
	"fmt"

	"github.com/gabe565/gh-profile/internal/profile"
	"github.com/gabe565/gh-profile/internal/util"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "switch [NAME]",
		Aliases: []string{"activate", "active", "sw", "s"},
		Short:   "Switch active profile",
		RunE:    run,
	}
	flagLocalDir(cmd)
	return cmd
}

func run(cmd *cobra.Command, args []string) (err error) {
	var p profile.Profile
	if len(args) > 0 {
		if args[0] == "-" {
			p, err = profile.GetPrevious()
			if err != nil {
				return err
			}
		} else {
			p = profile.New(args[0])
		}
	} else {
		if p, err = profile.Select("Choose a profile to activate"); err != nil {
			return err
		}
	}

	if inLocalDir {
		err = p.ActivateLocally(false)
	} else {
		err = p.ActivateGlobally(false)
	}
	if err != nil {
		if errors.Is(err, profile.ErrActive) {
			fmt.Println("ℹ️️ ", util.UpperFirst(err.Error()))
			return nil
		}
		return err
	}

	return nil
}
