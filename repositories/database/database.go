package database

import (
	"context"

	"github.com/barahouei/clean-architecture-telegram-bot/configs"
	"github.com/barahouei/clean-architecture-telegram-bot/models"
	"github.com/barahouei/clean-architecture-telegram-bot/pkg/logger"
	"github.com/barahouei/clean-architecture-telegram-bot/repositories"
	"github.com/barahouei/clean-architecture-telegram-bot/repositories/mongodb"
	"github.com/barahouei/clean-architecture-telegram-bot/repositories/mysql"
	"github.com/barahouei/clean-architecture-telegram-bot/repositories/postgres"
)

func NewDatabase(ctx context.Context, cfg *configs.Config, logger logger.Logger) (repositories.DB, error) {
	var db repositories.DB
	var err error

	switch cfg.App.Driver {
	case models.PostgreSQL.String():
		db, err = postgres.New(ctx, cfg.Postgres, logger)
		if err != nil {
			return nil, err
		}
	case models.MySQL.String():
		db, err = mysql.New(ctx, cfg.MySQL, logger)
		if err != nil {
			return nil, err
		}
	case models.MongoDB.String():
		db, err = mongodb.New(ctx, cfg.MongoDB, logger)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
