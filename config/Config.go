package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	Port        string `envconfig:"SERVER_PORT"`
	LibraryPath string `envconfig:"LIB_PATH"`
}

func GetConfig() Config {
	var cfg Config
	err := envconfig.Process("", &cfg)

	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
