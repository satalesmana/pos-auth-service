package config

import (
	"fmt"
	"go-grpc-auth-svc/pkg/models"

	"github.com/spf13/viper"
)

func LoadConfig() (config models.App, err error) {
	var AppConfig models.App

	viper.AddConfigPath("./files/config")
	viper.SetConfigType("yaml")
	viper.SetConfigName("default")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("%v", err)
	}

	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		fmt.Printf("unable to decode into config struct, %v", err)
	}

	return AppConfig, nil
}
