package config

import (
	"fmt"
	"github.com/spf13/viper"
	"sync"
)

var appConfigInstance *AppConfig
var loadErr error
var once sync.Once

func GetAppConfig() (*AppConfig, error) {
	once.Do(func() {
		appConfigInstance, loadErr = load("./config")
	})

	return appConfigInstance, loadErr
}

func setUpViper(configFolder string) {
	viper.AutomaticEnv()

	viper.SetDefault("APP_ENV", "development")
	appEnv := viper.GetString("APP_ENV")

	viper.AddConfigPath(configFolder)
	viper.SetConfigType("json")
	viper.SetConfigName(fmt.Sprintf("%s.%s", "config", appEnv))
}

func load(configFolder string) (*AppConfig, error) {
	setUpViper(configFolder)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	appConfig := getDefaultConfig()

	if err := viper.Unmarshal(appConfig); err != nil {
		return nil, err
	}

	return appConfig, nil
}
