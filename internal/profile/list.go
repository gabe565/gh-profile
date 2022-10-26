package profile

import (
	"errors"
	"os"
)

var ErrNoneFound = errors.New("no profiles found")

func List() ([]Profile, error) {
	files, err := os.ReadDir(ConfigDir())
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = ErrNoneFound
		}
		return []Profile{}, err
	}

	profiles := make([]Profile, 0, len(files))
	for _, file := range files {
		if file.IsDir() {
			profiles = append(profiles, Profile{Name: file.Name()})
		}
	}

	if len(profiles) == 0 {
		return []Profile{}, ErrNoneFound
	}

	return profiles, nil
}
