package profile

import (
	"errors"
	"github.com/gabe565/gh-profile/internal/github"
)

var ErrNoActive = errors.New("no profile is active")

type GetActiveFilter uint8

const (
	GetActiveDynamic GetActiveFilter = iota
	GetActiveGlobal
	GetActiveLocal
)

func GetActive(filters ...GetActiveFilter) (Profile, error) {
	var filter GetActiveFilter
	if len(filters) > 0 {
		filter = filters[0]
	}
	if filter == GetActiveDynamic {
		if github.ConfigDirOverridden() {
			filter = GetActiveLocal
		} else {
			filter = GetActiveGlobal
		}
	}

	profiles, err := List()
	if err != nil {
		return Profile{}, err
	}

	for _, profile := range profiles {
		switch filter {
		case GetActiveGlobal:
			if profile.Status().Global {
				return profile, nil
			}
		case GetActiveLocal:
			if profile.Status().Local {
				return profile, nil
			}
		}
	}

	return Profile{}, ErrNoActive
}
