package config

import (
	"github.com/caarlos0/env/v6"
	log "github.com/sirupsen/logrus"
)

type Configuration struct {
	AppName    string `env:"APP_NAME" envDefault:"blueprint-roa-golang"`
	Port       int    `env:"PORT" envDefault:"3040"`
	HttpPort   int    `env:"HTTP_PORT" envDefault:"3041"`
	Env        string `env:"ENV" envDefault:"localhost"`
	DbDriver   string `env:"DB_DRIVER" envDefault:"postgres"`
	DbHost     string `env:"DB_HOST" envDefault:"127.0.0.1"`
	DbPort     string `env:"DB_PORT" envDefault:"5432"`
	DbUser     string `env:"DB_USER" envDefault:"postgres"`
	DbName     string `env:"DB_NAME" envDefault:"users"`
	DbPassword string `env:"DB_PASSWORD" envDefault:"postgres"`
	Jaeger     Jaeger
}

func NewConfiguration() Configuration {
	config := Configuration{}

	if err := env.Parse(&config); err != nil {
		log.Errorf("%+v\n", err)
	}

	return config
}

// This should be in an env file in production (24 digits)
const MySecret string = "kgmj12@nJuh8#od6gijr4(&2"
