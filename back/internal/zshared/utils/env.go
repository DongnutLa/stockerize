package utils

import (
	"errors"
	"fmt"

	"sync"

	"github.com/spf13/viper"
)

var (
	config = make(map[string]string)
	once   sync.Once
)

func LoadConfig() (map[string]string, error) {
	once.Do(func() {
		// Initialize Viper
		viper.AddConfigPath(".")
		viper.SetConfigFile(".env")
		viper.SetConfigType("env")

		// Read the environment variables
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("Error reading env file. %+v", err)
			return
		}

		// Iterate through the keys in the configuration
		for _, key := range viper.AllKeys() {
			value := viper.GetString(key)
			config[key] = value
		}
	})

	if config == nil {
		return nil, errors.New("failed to load env variables")
	}

	return config, nil
}

func GetConfig(key string) string {
	return config[key]
}
