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
	profiles, err := profile.List()
	if err != nil {
		return err
	}

	for _, p := range profiles {
		if p.Status().IsActive() {
			fmt.Println(p.Name)
			return nil
		}
	}

	fmt.Println("none")
	return nil
}
