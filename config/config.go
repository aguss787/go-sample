package config

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/dig"
	"log"
	"os"
	"strings"
)

type Config struct {
	Sample   string         `mapstructure:"sample"`
	Postgres PostgresConfig `mapstructure:"postgres"`
}

type PostgresConfig struct {
	DatabaseUrl string `mapstructure:"database_url"`
}

func readConfig() Config {
	runEnv, exists := os.LookupEnv("GLINTS_RUN_ENV")
	if !exists {
		runEnv = "development"
	}

	configDir, exists := os.LookupEnv("GLINTS_CONFIG_DIR")
	if !exists {
		configDir = "var"
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	viper.SetEnvPrefix("glints")

	viper.SetConfigFile(fmt.Sprintf("%s/%s.yaml", configDir, runEnv))
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Unable to read config: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to unmarshal config: %v", err)
	}

	return config
}

func Register(container *dig.Container) (*dig.Container, error) {
	if err := container.Provide(readConfig); err != nil {
		return nil, err
	}

	return container, nil
}
