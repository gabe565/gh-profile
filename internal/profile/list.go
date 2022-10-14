package profile

import (
	"github.com/gabe565/gh-profile/internal/github"
	"os"
	"path/filepath"
)

func List() ([]Profile, error) {
	conf := github.ConfigDir()

	files, err := os.ReadDir(filepath.Join(conf, "profiles"))
	if err != nil {
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
