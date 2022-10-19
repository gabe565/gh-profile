package _switch

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var inLocalDir bool

func flagLocalDir(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&inLocalDir, "local-dir", "l", false, "Enables the profile for the current directory using direnv.")
	if err := viper.BindPFlag("local-dir", cmd.Flags().Lookup("local-dir")); err != nil {
		panic(err)
	}
}
