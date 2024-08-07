package config

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/viper"
)

func Setup(path string) {
	fileName := os.Getenv("HOST")
	if fileName == "" {
		fileName = "local"
	}

	viper.AddConfigPath(path)
	viper.SetConfigName(fileName)
	viper.SetConfigType("yaml")

	// Read the config file
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}

	// Here, we want viper to update the key values it holds
	// for every value that start with $, such as $HELLO
	// update the value by fetching the value from the environment of key HELLO
	for key := range viper.AllSettings() {
		replaceEnvKeys(key)
	}

	slog.Info("Using config from", "file", viper.ConfigFileUsed())
	slog.Info("Config loaded successfully")
}

func replaceEnvKeys(key string) {
	value := viper.Get(key)
	switch value := value.(type) {
	case string:
		if strings.HasPrefix(value, "$") {
			viper.Set(key, os.Getenv(strings.TrimPrefix(value, "$")))
		}
	case map[string]interface{}:
		for k := range value {
			newKey := fmt.Sprintf("%s.%s", key, k)
			replaceEnvKeys(newKey)
		}
	}

}
