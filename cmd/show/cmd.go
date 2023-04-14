package show

import (
	"fmt"

	"github.com/gabe565/gh-profile/internal/profile"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Shows the active profile name",
		RunE:  run,
	}
}

func run(cmd *cobra.Command, args []string) (err error) {
	p, _ := profile.GetActive()
	fmt.Println(p.Name)

	return nil
}
