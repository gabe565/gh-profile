package github

import (
	"path/filepath"
)

const HostsFilename = "hosts.yml"

func HostsPath() string {
	return filepath.Join(ConfigDir(), HostsFilename)
}
