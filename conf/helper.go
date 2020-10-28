package conf

import "github.com/spf13/viper"

func GetFileApiPort() string {
	return viper.GetString("api.fileApi")
}

func GetPingApiPort() string {
	return viper.GetString("api.pingApi")
}

func GetFsRootDir() string {
	return viper.GetString("fs.rootDir")
}

func GetLogCategory() string {
	return viper.GetString("log.category")
}