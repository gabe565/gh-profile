package profile

import (
	"errors"
	"fmt"
	"github.com/gabe565/gh-profile/internal/github"
	"github.com/gabe565/gh-profile/internal/util"
	"os"
	"os/exec"
	"path/filepath"
)

func New(name string) Profile {
	return Profile{
		Name: filepath.Base(name),
	}
}

var ErrExist = errors.New("profile already exists")

var ErrNotExist = errors.New("profile does not exist")

var ErrActive = errors.New("profile already active")

var ErrNameUnchanged = errors.New("name unchanged")

var ConfigDirName = "profiles"

func ConfigDir() string {
	return filepath.Join(github.RootConfigDir(), ConfigDirName)
}

type Profile struct {
	Name string
}

func (p Profile) Path() string {
	return filepath.Join(ConfigDir(), filepath.Base(p.Name))
}

func (p Profile) HostsPath() string {
	return filepath.Join(p.Path(), github.HostsFilename)
}

func (p Profile) ConfigPath() string {
	return filepath.Join(p.Path(), github.ConfigFilename)
}

func (p Profile) Exists() bool {
	if _, err := os.Stat(p.Path()); err != nil {
		return false
	}
	return true
}

func (p Profile) Create() error {
	if p.Exists() {
		return fmt.Errorf("%w: %s", ErrExist, p.Name)
	}

	fmt.Println("✨ Creating profile:", p.Name)

	// Create profile dir
	if err := os.MkdirAll(p.Path(), 0755); err != nil {
		return err
	}

	// Create profile hosts
	if err := util.CopyFile(github.RootHostsPath(), p.HostsPath()); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	// Create profile config
	if err := util.CopyFile(github.RootConfigPath(), p.ConfigPath()); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	return nil
}

func (p Profile) Remove() error {
	if !p.Exists() {
		return fmt.Errorf("%w: %s", ErrNotExist, p.Name)
	}

	fmt.Println("🔥 Removing profile:", p.Name)

	return os.RemoveAll(p.Path())
}

func (p Profile) ActivateLocally(force bool) error {
	if !p.Exists() {
		return fmt.Errorf("%w: %s", ErrNotExist, p.Name)
	}

	fmt.Println("🔧 Activating local dir profile:", p.Name)

	if _, err := exec.LookPath("direnv"); err != nil {
		fmt.Println("⚠️  direnv not found. To use local dir profiles, please see https://direnv.net")
	}

	if p.IsActiveLocally() && !force {
		return fmt.Errorf("%w: %s", ErrActive, p.Name)
	}

	// Copy config to profile if not exist
	if _, err := os.Stat(p.ConfigPath()); errors.Is(err, os.ErrNotExist) {
		if err := util.CopyFile(github.RootConfigPath(), p.ConfigPath()); err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return err
			}
		}
	}

	f, err := os.OpenFile(".envrc", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
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

func (p Profile) ActivateGlobally(force bool) error {
	if !p.Exists() {
		return fmt.Errorf("%w: %s", ErrNotExist, p.Name)
	}

	if p.IsActiveGlobally() && !force {
		return fmt.Errorf("%w: %s", ErrActive, p.Name)
	}

	if github.ConfigDirOverridden() {
		fmt.Println("ℹ️  Found local dir profile, but global change was requested")
	}

	fmt.Println("🔧 Activating global profile:", p.Name)

	// Remove existing hosts
	if err := os.Remove(github.RootHostsPath()); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	// Symlink profile hosts
	if err := os.Symlink(p.HostsPath(), github.RootHostsPath()); err != nil {
		return err
	}

	// Copy config to profile if not exist
	if _, err := os.Stat(p.ConfigPath()); errors.Is(err, os.ErrNotExist) {
		if err := util.CopyFile(github.RootConfigPath(), p.ConfigPath()); err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return err
			}
		}
	}

	// Remove existing config
	if err := os.Remove(github.RootConfigPath()); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	// Symlink profile config
	if err := os.Symlink(p.ConfigPath(), github.RootConfigPath()); err != nil {
		return err
	}

	return nil
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

func (p Profile) Rename(to string) error {
	if !p.Exists() {
		return fmt.Errorf("%w: %s", ErrNotExist, p.Name)
	}

	if to == p.Name {
		return fmt.Errorf("%w: %s to %s", ErrNameUnchanged, p.Name, to)
	}

	fmt.Println("🚚 Renaming profile:", p.Name, "to", to)

	wasActive := p.IsActiveGlobally()

	oldPath := p.Path()
	p.Name = to
	if err := os.Rename(oldPath, p.Path()); err != nil {
		return err
	}

	if wasActive {
		return p.ActivateGlobally(false)
	}

	return nil
}
