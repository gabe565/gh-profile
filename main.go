package main

import (
	"fmt"
	"github.com/gabe565/gh-profile/cmd"
	"github.com/gabe565/gh-profile/internal/util"
	"os"
)

//go:generate go run ./internal/cmd/docs

func main() {
	if err := cmd.Command.Execute(); err != nil {
		fmt.Println("🚫", util.UpperFirst(err.Error()))
		os.Exit(1)
	}
}
