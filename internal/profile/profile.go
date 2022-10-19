package profile

import (
	"errors"
	"fmt"
	"github.com/gabe565/gh-profile/internal/github"
	"os"
	"os/exec"
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
	conf := github.RootConfigDir()
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

func (p Profile) Remove() error {
	fmt.Println("🔥 Removing profile", p.Name)

	if !p.Exists() {
		return ErrProfileNotExist
	}

	return os.RemoveAll(p.Path())
}

var ErrProfileNotExist = errors.New("profile does not exist")

var ErrProfileActive = errors.New("profile already active")

func (p Profile) ActivateLocally() error {
	if !p.Exists() {
		return ErrProfileNotExist
	}

	fmt.Println("🔧 Activating local dir profile", p.Name)

	if _, err := exec.LookPath("direnv"); err != nil {
		fmt.Println("⚠️  direnv not found. To use local dir profiles, please see https://direnv.net")
	}

	if p.IsActiveLocally() {
		return ErrProfileActive
	}

	f, err := os.OpenFile(".envrc", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	if _, err := f.WriteString("export GH_CONFIG_DIR='" + p.Path() + "'\n"); err != nil {
		return err
	}

	return f.Close()
}

func (p Profile) ActivateGlobally() error {
	if !p.Exists() {
		return ErrProfileNotExist
	}

	fmt.Println("🔧 Activating global profile", p.Name)

	if p.IsActiveGlobally() {
		return ErrProfileActive
	}

	// Remove existing hosts config
	if err := os.Remove(github.RootHostsPath()); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	// Hardlink profile hosts config
	return os.Symlink(p.HostsPath(), github.RootHostsPath())
}

func (p Profile) Status() ActiveStatus {
	if p.IsActiveLocally() {
		return StatusLocal
	}
	if p.IsActiveGlobally() {
		return StatusGlobal
	}
	return StatusInactive
}

func (p Profile) IsActiveGlobally() bool {
	target, err := filepath.EvalSymlinks(github.RootHostsPath())
	if err == nil && target == p.HostsPath() {
		return true
	}
	return false
}

func (p Profile) IsActiveLocally() bool {
	if github.ConfigDirOverridden() {
		overrideName := filepath.Base(github.ConfigDir())
		if overrideName == p.Name {
			return true
		}
	}
	return false
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

	wasActive := p.IsActiveGlobally()

	oldPath := p.Path()
	p.Name = to
	if err := os.Rename(oldPath, p.Path()); err != nil {
		return err
	}

	if wasActive {
		return p.ActivateGlobally()
	}

	return nil
}
