package main

import (
	"github.com/gabe565/gh-profile/cmd"
	"os"
)

func main() {
	if err := cmd.Command.Execute(); err != nil {
		os.Exit(1)
	}
}
