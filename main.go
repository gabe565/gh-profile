package main

import (
	"github.com/gabe565/gh-profile/cmd"
	"os"
)

//go:generate go run ./internal/cmd/docs

func main() {
	if err := cmd.Command.Execute(); err != nil {
		os.Exit(1)
	}
}
