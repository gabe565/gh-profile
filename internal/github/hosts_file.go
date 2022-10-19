package github

import (
	"path/filepath"
)

const HostsFilename = "hosts.yml"

func HostsPath() string {
	return filepath.Join(ConfigDir(), HostsFilename)
}

func RootHostsPath() string {
	return filepath.Join(RootConfigDir(), HostsFilename)
}
