package profile

import (
	"errors"
	"fmt"
	"github.com/gabe565/gh-profile/internal/github"
	"github.com/gabe565/gh-profile/internal/util"
	"os"
	"path/filepath"
)

func ExistingToDefault() error {
	conf := github.ConfigDir()

	isLink, err := util.IsLink(filepath.Join(conf, "hosts.yml"))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if !isLink {
		fmt.Println("Creating default profile")
		p := Profile{Name: "default"}

		if err := os.MkdirAll(p.Path(), 0755); err != nil {
			return err
		}

		if err := os.Rename(github.HostsPath(), p.HostsPath()); err != nil {
			return err
		}

		return p.Activate()
	}

	return nil
}
