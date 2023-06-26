// Package configs implement configurations that the application needs.
package configs

import (
	"github.com/barahouei/clean-architecture-telegram-bot/models"
	"github.com/barahouei/clean-architecture-telegram-bot/pkg/logger"
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		App      App
		Postgres Postgres
		MongoDB  MongoDB
	}

	App struct {
		Name  string `env:"BOT_NAME,required"`
		Token string `env:"BOT_APITOKEN,required"`
	}

	Postgres struct {
		Host     string `env:"BOT_POSTGRES_HOST,required"`
		Port     string `env:"BOT_POSTGRES_PORT,required"`
		Username string `env:"BOT_POSTGRES_USERNAME,required"`
		Password string `env:"BOT_POSTGRES_PASSWORD,required"`
		DBName   string `env:"BOT_POSTGRES_DBNAME,required"`
		SSL      string `env:"BOT_POSTGRES_SSL_MODE,required"`
	}

	MongoDB struct {
		Protocol string `env:"BOT_MONGODB_PROTOCOL,required"`
		Username string `env:"BOT_MONGODB_USERNAME,required"`
		Password string `env:"BOT_MONGODB_PASSWORD,required"`
		Host     string `env:"BOT_MONGODB_HOST,required"`
		Port     string `env:"BOT_MONGODB_PORT,required"`
		DBName   string `env:"BOT_MONGODB_DBNAME,required"`
	}
)

// New returns the config, if can't open .env file or parse environment variables it returns an error.
//
// make sure to delete the .env file and pass production mode in production.
func New(logger logger.Logger, mode models.Mode) (*Config, error) {
	if mode == models.Development {
		err := godotenv.Load("configs/.env")
		if err != nil {
			logger.Error(err)

			return nil, err
		}
	}

	cfg := Config{}

	err := env.Parse(&cfg)
	if err != nil {
		logger.Error(err)

		return nil, err
	}

	return &cfg, nil
}
