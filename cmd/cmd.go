package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"syscall"
)

var Command = &cobra.Command{
	Use:  "profile name",
	RunE: run,
}

func run(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("profile name is required")
	}
	profile := args[0]

	conf, err := getConfigDir()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(conf, "accounts"), 0755); err != nil {
		return err
	}

	if isLink, err := isLink(filepath.Join(conf, "hosts.yml")); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return err
		}
	} else if !isLink {
		fmt.Println(`Naming existing profile to default`)
		if err := newDefaultProfile(); err != nil {
			return err
		}
	}

	fmt.Println("Switching to profile", profile)
	if err := switchProfile(profile); err != nil {
		return err
	}

	return nil
}

func getConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".config", "gh"), nil
}

func isLink(path string) (bool, error) {
	fi, err := os.Lstat(path)
	if err != nil {
		return false, err
	}

	if fi.Mode()&os.ModeSymlink != 0 {
		return true, nil
	}

	s, ok := fi.Sys().(*syscall.Stat_t)
	if !ok {
		return false, errors.New("cannot convert stat value to syscall.Stat_t")
	}

	return s.Nlink > 1, nil
}

func newDefaultProfile() error {
	conf, err := getConfigDir()
	if err != nil {
		return err
	}

	profilePath := filepath.Join(conf, "accounts", "default")
	if err := os.MkdirAll(profilePath, 0755); err != nil {
		return err
	}

	ghHosts := filepath.Join(conf, "hosts.yml")
	profileHosts := filepath.Join(profilePath, "hosts.yml")
	if err := os.Rename(ghHosts, profileHosts); err != nil {
		return err
	}

	if err := os.Link(profileHosts, ghHosts); err != nil {
		return err
	}

	return nil
}

func switchProfile(name string) error {
	conf, err := getConfigDir()
	if err != nil {
		return err
	}

	profilePath := filepath.Join(conf, "accounts", filepath.Join("/", name))
	if err := os.MkdirAll(profilePath, 0755); err != nil {
		return err
	}

	ghHosts := filepath.Join(conf, "hosts.yml")
	profileHosts := filepath.Join(profilePath, "hosts.yml")
	if err := os.Remove(ghHosts); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	f, err := os.OpenFile(filepath.Join(profilePath, "hosts.yml"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	if err := os.Link(profileHosts, ghHosts); err != nil {
		return err
	}

	return nil
}
