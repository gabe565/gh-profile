package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gabe565/gh-profile/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	output := "./docs"

	if err := os.RemoveAll(output); err != nil {
		log.Fatal(fmt.Errorf("failed to remove existing dir: %w", err))
	}

	if err := os.MkdirAll(output, 0o755); err != nil {
		log.Fatal(fmt.Errorf("failed to mkdir: %w", err))
	}

	cmd.DefaultConfigDir = "$HOME/.config/gh"

	rootCmd := cmd.New("latest", "")
	rootCmd.InitDefaultVersionFlag()

	if err := doc.GenMarkdownTree(rootCmd, output); err != nil {
		log.Fatal(fmt.Errorf("failed to generate markdown: %w", err))
	}
}
