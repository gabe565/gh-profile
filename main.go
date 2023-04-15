package main

import (
	"log"

	"github.com/gabe565/gh-profile/cmd"
	"github.com/gabe565/gh-profile/internal/util"
)

var (
	version = "next"
	commit  = ""
)

func main() {
	rootCmd := cmd.New(version, commit)
	if err := rootCmd.Execute(); err != nil {
		l := log.New(rootCmd.ErrOrStderr(), "", 0)
		l.Fatalln("🚫", util.UpperFirst(err.Error()))
	}
}
