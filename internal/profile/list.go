package profile

import (
	"errors"
	"github.com/gabe565/gh-profile/internal/github"
	"os"
	"path/filepath"
)

var ErrNoProfiles = errors.New("no profiles found")

func List() ([]Profile, error) {
	conf := github.RootConfigDir()

	files, err := os.ReadDir(filepath.Join(conf, "profiles"))
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = ErrNoProfiles
		}
		return []Profile{}, err
	}

	profiles := make([]Profile, 0, len(files))
	for _, file := range files {
		if file.IsDir() {
			profiles = append(profiles, Profile{Name: file.Name()})
		}
	}

	return profiles, nil
}
