package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func ReadConfigFiles(path, file string) {

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)

	viper.AddConfigPath(path)
	viper.AddConfigPath(exPath)
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName("common")
	viper.ReadInConfig()

	if file != "" {
		file = path + file
		var _, err = os.Stat(file)
		if !os.IsNotExist(err) {
			viper.SetConfigFile(file)
			viper.MergeInConfig()
		} else {
			panic(fmt.Sprintf("Config file \"%s\" was not found", file))
		}
	}
}
