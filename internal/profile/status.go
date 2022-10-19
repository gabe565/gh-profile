package profile

import "github.com/gabe565/gh-profile/internal/github"

const (
	StatusInactive ActiveStatus = iota
	StatusGlobal
	StatusLocal
)

type ActiveStatus uint8

func (s ActiveStatus) IsActive() bool {
	if github.ConfigDirOverridden() {
		return s == StatusLocal
	} else {
		return s == StatusGlobal
	}
}

func (s ActiveStatus) IsAnyActive() bool {
	return s == StatusGlobal || s == StatusLocal
}
