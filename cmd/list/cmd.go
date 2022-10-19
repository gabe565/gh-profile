package list

import (
	"fmt"
	"github.com/gabe565/gh-profile/internal/profile"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls", "l"},
		Short:   "Lists all profiles",
		RunE:    run,
	}
}

func run(cmd *cobra.Command, args []string) (err error) {
	profiles, err := profile.List()
	if err != nil {
		return err
	}

	for _, p := range profiles {
		if p.Status().IsActive() {
			fmt.Println(text.FgGreen.Sprint("✓"), text.Bold.Sprint(p.Name))
		} else {
			fmt.Println(" ", p.Name)
		}
	}

	return nil
}
