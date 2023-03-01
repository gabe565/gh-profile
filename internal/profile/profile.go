package profile

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gabe565/gh-profile/internal/github"
	"github.com/gabe565/gh-profile/internal/util"
)

func New(name string) Profile {
	return Profile{
		Name: filepath.Base(name),
	}
}

var (
	ErrExist         = errors.New("profile already exists")
	ErrNotExist      = errors.New("profile does not exist")
	ErrActive        = errors.New("profile already active")
	ErrNameUnchanged = errors.New("name unchanged")
	ConfigDirName    = "profiles"
)

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

	var first bool
	if _, err := List(); err != nil {
		if !errors.Is(err, ErrNoneFound) {
			return err
		}
		first = true
	}

	// Create profile dir
	if err := os.MkdirAll(p.Path(), 0755); err != nil {
		return err
	}

	// Create profile hosts
	if first {
		fmt.Println("ℹ️️  Copying existing account into profile")
		if err := util.CopyFile(github.RootHostsPath(), p.HostsPath()); err != nil && !errors.Is(err, os.ErrNotExist) {
			return err
		}
	} else {
		f, err := os.Create(p.HostsPath())
		if err != nil {
			return err
		}
		if err := f.Close(); err != nil {
			return err
		}
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

	if _, err := exec.LookPath("direnv"); err != nil {
		fmt.Println("⚠️  direnv not found. To use local dir profiles, please see https://direnv.net")
	}

	if p.IsActiveLocally() && !force {
		return fmt.Errorf("%w: %s", ErrActive, p.Name)
	}

	fmt.Println("🔧 Activating local dir profile:", p.Name)

	// Copy config to profile if not exist
	if _, err := os.Stat(p.ConfigPath()); errors.Is(err, os.ErrNotExist) {
		if err := util.CopyFile(github.RootConfigPath(), p.ConfigPath()); err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return err
			}
		}
	}

	f, err := os.OpenFile(".envrc", os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	envPrefix := "export GH_CONFIG_DIR="
	envLine := envPrefix + `"` + util.ReplaceEnvsInPath(p.Path()) + `"`

	var buf bytes.Buffer
	scanner := bufio.NewScanner(f)
	var found bool
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, envPrefix) {
			if found {
				continue
			} else {
				line = envLine
				found = true
			}
		}
		if _, err := buf.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	if scanner.Err() != nil {
		return scanner.Err()
	}

	if !found {
		buf.WriteString(envLine + "\n")
	}

	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return err
	}
	if err := f.Truncate(0); err != nil {
		return err
	}

	if _, err := io.Copy(f, &buf); err != nil {
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

	prevProfile, _ := GetActive(GetActiveGlobal)

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

	if prevProfile.Name != "" {
		if err := prevProfile.WritePrevious(); err != nil {
			fmt.Println("⚠️  Failed to store previous profile")
		}
	}

	exec.Command("gh", "auth", "setup-git")

	return nil
}

func (p Profile) Status() ActiveStatus {
	status := ActiveStatus{
		Global: p.IsActiveGlobally(),
		Local:  p.IsActiveLocally(),
	}
	return status
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
