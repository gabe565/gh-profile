package main

import (
	"fmt"
	"github.com/gabe565/gh-profile/cmd"
	"os"
)

//go:generate go run ./internal/cmd/docs

func main() {
	if err := cmd.Command.Execute(); err != nil {
		fmt.Println("🚫", err)
		os.Exit(1)
	}
}
