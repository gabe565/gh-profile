package profile

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gabe565/gh-profile/internal/github"
	"os"
	"path/filepath"
)

func (p Profile) WritePrevious() error {
	f, err := os.Create(filepath.Join(github.RootConfigDir(), ConfigDirName, "previous"))
	if err != nil {
		return err
	}

	if _, err := f.WriteString(p.Name); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}

var (
	ErrPreviousNotSet   = errors.New("previous profile not set")
	ErrPreviousNotExist = errors.New("previous profile does not exist")
)

func GetPrevious() (Profile, error) {
	contents, err := os.ReadFile(filepath.Join(github.RootConfigDir(), ConfigDirName, "previous"))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return Profile{}, ErrPreviousNotSet
		}
		return Profile{}, err
	}

	name := string(bytes.TrimSpace(contents))

	profiles, err := List()
	if err != nil {
		return Profile{}, err
	}

	for _, profile := range profiles {
		if profile.Name == string(name) {
			return profile, nil
		}
	}

	return Profile{}, fmt.Errorf("%w: %s", ErrPreviousNotExist, string(name))
}
