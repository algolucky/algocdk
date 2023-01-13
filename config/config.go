package config

import (
	"github.com/spf13/viper"
)

// C exposes the config type as a variable
var C config

var (
	defaultConfigName = "algocdk"
	defaultConfigPath = "."
	defaultConfigType = "yaml"
)

// InitConfig takes a path to a viper config file and tries to load it
func InitConfig(cfgFile string) {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(defaultConfigPath)
		viper.SetConfigName(defaultConfigName)
		viper.SetConfigType(defaultConfigType)
	}

	viper.ReadInConfig()
}

// Unmarshal load the viper config into the variable C
func Unmarshal() (err error) {
	err = viper.GetViper().Unmarshal(&C)
	return
}
