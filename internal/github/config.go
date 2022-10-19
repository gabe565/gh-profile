package github

import "github.com/spf13/viper"

const ConfigDirKey = "config-dir"

func ConfigDir() string {
	return viper.GetString(ConfigDirKey)
}

func SetConfigDir(path string) {
	viper.Set(ConfigDirKey, path)
}
