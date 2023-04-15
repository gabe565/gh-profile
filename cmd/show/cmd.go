package show

import (
	"fmt"

	"github.com/gabe565/gh-profile/internal/profile"
	"github.com/gabe565/gh-profile/internal/util"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	return &cobra.Command{
		Use:     "show",
		Short:   "Shows the active profile name",
		GroupID: "read",
		RunE:    run,

		ValidArgsFunction: util.ShellCompDisable,
	}
}

func run(cmd *cobra.Command, args []string) (err error) {
	p, err := profile.GetActive()
	if err != nil {
		return err
	}

	fmt.Println(p.Name)
	return nil
}
