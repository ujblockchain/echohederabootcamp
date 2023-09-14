package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(env string) {
	//
	var err error
	config = viper.New()

	//set file format
	config.SetConfigType("yaml")

	//set the config file name
	config.SetConfigName(env)

	// find file folder
	config.AddConfigPath("config/")

	//try to read file to see if error exist
	err = config.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("error file not found %w", err))
	}

}

func GetConfig(env string) *viper.Viper {
	Init(env)

	return config
}
