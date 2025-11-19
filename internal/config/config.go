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

/*

package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	//environment
	Environment string `mapstructure:"ENV"`
	// server
	AppPort uint16 `mapstructure:"APP_PORT"`
	// database
	MongoURI       string `mapstructure:"MONGO_URI"`
	DBName         string `mapstructure:"DB_NAME"`
	DBQueryTimeout uint16 `mapstructure:"DB_QUERY_TIMEOUT_SEC"`
	// Token
	JWTSECRET   string `mapstructure:"JWT_SECRET"`
	TokenExpiry uint16 `mapstructure:"TOKEN_EXPIRY_DAYS"`
}

var Environ = Env{}

func NewEnv(filename string, override bool) *Env {
	viper.SetConfigFile(filename)

	if override {
		viper.AutomaticEnv()
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading environment file", err)
	}

	err = viper.Unmarshal(&Environ)
	if err != nil {
		log.Fatal("Error loading environment file", err)
	}

	return &Environ
}
*/
