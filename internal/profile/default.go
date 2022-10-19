package profile

import (
	"errors"
	"fmt"
	"github.com/gabe565/gh-profile/internal/github"
	"github.com/gabe565/gh-profile/internal/util"
	"os"
)

func ExistingToDefault() error {
	isLink, err := util.IsLink(github.RootHostsPath())
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if !isLink {
		fmt.Println("✨ Creating default profile")
		p := Profile{Name: "default"}

		if err := os.MkdirAll(p.Path(), 0755); err != nil {
			return err
		}

		if err := os.Rename(github.RootHostsPath(), p.HostsPath()); err != nil {
			return err
		}

		if err := os.Rename(github.RootConfigPath(), p.ConfigPath()); err != nil {
			return err
		}
	}

	return nil
}
