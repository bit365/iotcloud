package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func LoadConfig() *viper.Viper {
	config := viper.New()

	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath("./config")

	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("Error reading default config file: %s", err)
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}
	config.SetConfigName("config." + env)

	if err := config.MergeInConfig(); err != nil {
		log.Printf("No environment-specific config file found for '%s', using defaults", env)
	}

	return config
}
