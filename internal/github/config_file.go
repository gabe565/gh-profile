package github

import "path/filepath"

const ConfigFilename = "config.yml"

func ConfigPath() string {
	return filepath.Join(ConfigDir(), ConfigFilename)
}

func RootConfigPath() string {
	return filepath.Join(RootConfigDir(), ConfigFilename)
}
