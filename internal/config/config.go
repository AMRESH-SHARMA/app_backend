package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func Load() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev" // default to dev
	}
	viper.SetConfigName("config." + env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // or your config directory

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
