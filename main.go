package main

import (
	"fmt"
	"github.com/gabe565/gh-profile/cmd"
	"github.com/gabe565/gh-profile/internal/util"
	"os"
)

func main() {
	if err := cmd.New().Execute(); err != nil {
		fmt.Println("🚫", util.UpperFirst(err.Error()))
		os.Exit(1)
	}
}
