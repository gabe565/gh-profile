package profile

import (
	"errors"
	"fmt"
	"github.com/gabe565/gh-profile/internal/github"
	"os"
	"path/filepath"
)

func New(name string) Profile {
	return Profile{
		Name: filepath.Base(name),
	}
}

type Profile struct {
	Name string
}

func (p Profile) Path() string {
	conf := github.ConfigDir()
	return filepath.Join(conf, "accounts", filepath.Join("/", p.Name))
}

func (p Profile) HostsPath() string {
	return filepath.Join(p.Path(), github.HostsFilename)
}

func (p Profile) Create() error {
	// Create profile dir
	if err := os.MkdirAll(p.Path(), 0755); err != nil {
		return err
	}

	// Create profile hosts config
	f, err := os.OpenFile(filepath.Join(p.Path(), "hosts.yml"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	return f.Close()
}

func (p Profile) Activate() error {
	if _, err := os.Stat(p.Path()); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}

		fmt.Println("Creating profile", p.Name)
		if err := p.Create(); err != nil {
			return err
		}
	}

	fmt.Println("Activating", p.Name)

	// Remove existing hosts config
	if err := os.Remove(github.HostsPath()); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	// Hardlink profile hosts config
	if err := os.Link(p.HostsPath(), github.HostsPath()); err != nil {
		return err
	}

	return nil
}
