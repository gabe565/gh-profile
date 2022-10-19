package github

import "github.com/spf13/viper"

const (
	ConfigDirKey     = "config-dir"
	RootConfigDirKey = "root-config-dir"
)

func ConfigDir() string {
	return viper.GetString(ConfigDirKey)
}

func SetConfigDir(path string) {
	viper.Set(ConfigDirKey, path)
}

func RootConfigDir() string {
	return viper.GetString(RootConfigDirKey)
}

func SetRootConfigDir(path string) {
	viper.Set(RootConfigDirKey, path)
}

func ConfigDirOverridden() bool {
	return ConfigDir() != RootConfigDir()
}
