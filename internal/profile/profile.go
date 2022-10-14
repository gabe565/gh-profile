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
	return filepath.Join(conf, "profiles", filepath.Base(p.Name))
}

func (p Profile) HostsPath() string {
	return filepath.Join(p.Path(), github.HostsFilename)
}

func (p Profile) Exists() bool {
	if _, err := os.Stat(p.Path()); err != nil {
		return false
	}
	return true
}

var ErrProfileExist = errors.New("profile already exists")

func (p Profile) Create() error {
	fmt.Println("✨ Creating profile", p.Name)

	if p.Exists() {
		return ErrProfileExist
	}

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

func (p Profile) Delete() error {
	fmt.Println("🔥 Deleting profile", p.Name)
	return os.RemoveAll(p.Path())
}

var ErrProfileNotExist = errors.New("profile does not exist")

var ErrProfileActive = errors.New("profile already active")

func (p Profile) Activate() error {
	fmt.Println("🔧 Activating profile", p.Name)

	if !p.Exists() {
		return ErrProfileNotExist
	}

	if p.IsActive() {
		return ErrProfileActive
	}

	// Remove existing hosts config
	if err := os.Remove(github.HostsPath()); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	// Hardlink profile hosts config
	if err := os.Symlink(p.HostsPath(), github.HostsPath()); err != nil {
		return err
	}

	return nil
}

func (p Profile) IsActive() bool {
	target, err := filepath.EvalSymlinks(github.HostsPath())
	if err != nil {
		return false
	}

	return target == p.HostsPath()
}

var ErrSameName = errors.New("name unchanged")

func (p Profile) Rename(to string) error {
	fmt.Println("🚚 Renaming", p.Name, "to", to)

	if !p.Exists() {
		return ErrProfileNotExist
	}

	if to == p.Name {
		return ErrSameName
	}

	wasActive := p.IsActive()

	oldPath := p.Path()
	p.Name = to
	if err := os.Rename(oldPath, p.Path()); err != nil {
		return err
	}

	if wasActive {
		return p.Activate()
	}
	return nil
}
