package main

import (
	"fmt"
	"os"

	"github.com/gabe565/gh-profile/cmd"
	"github.com/gabe565/gh-profile/internal/util"
)

var (
	version = "next"
	commit  = ""
)

func main() {
	if err := cmd.New(version, commit).Execute(); err != nil {
		fmt.Println("🚫", util.UpperFirst(err.Error()))
		os.Exit(1)
	}
}
