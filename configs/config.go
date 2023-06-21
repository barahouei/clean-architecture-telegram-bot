// Package configs implement configurations that the application needs.
package configs

import (
	"fmt"

	"github.com/barahouei/clean-architecture-telegram-bot/pkg/logger"
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		App      App
		Postgres Postgres
	}

	App struct {
		Name  string `env:"BOT_NAME,required"`
		Token string `env:"BOT_APITOKEN,required"`
	}

	Postgres struct {
		Host     string `env:"BOT_POSTGRES_HOST,required"`
		Port     string `env:"BOT_POSTGRES_PORT,required"`
		User     string `env:"BOT_POSTGRES_USER,required"`
		Password string `env:"BOT_POSTGRES_PASSWORD,required"`
		DBName   string `env:"BOT_POSTGRES_DBNAME,required"`
		SSL      string `env:"BOT_POSTGRES_SSL_MODE,required"`
	}

	Mode uint
)

const (
	_ Mode = iota
	development
	production
)

// New returns the config, if can't open .env file or parse environment variables it returns an error.
//
// make sure to delete the .env file and pass production mode in production.
func New(logger logger.Logger, mode Mode) (*Config, error) {
	if mode == development {
		err := godotenv.Load("configs/.env")
		if err != nil {
			logger.Error(fmt.Sprintf("unable to load .env file: %e", err))

			return nil, err
		}
	}

	cfg := Config{}

	err := env.Parse(&cfg)
	if err != nil {
		logger.Error(fmt.Sprintf("unable to parse ennvironment variables: %e", err))

		return nil, err
	}

	return &cfg, nil
}
