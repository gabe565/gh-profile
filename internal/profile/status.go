package profile

import "github.com/gabe565/gh-profile/internal/github"

type ActiveStatus struct {
	Global bool
	Local  bool
}

func (s ActiveStatus) IsActive() bool {
	if github.ConfigDirOverridden() {
		return s.Local
	} else {
		return s.Global
	}
}

func (s ActiveStatus) IsAnyActive() bool {
	return s != ActiveStatus{}
}
