// Package mongodb implements mongodb database functionalities.
package mongodb

import (
	"context"
	"fmt"

	"github.com/barahouei/clean-architecture-telegram-bot/configs"
	"github.com/barahouei/clean-architecture-telegram-bot/pkg/logger"
	"github.com/barahouei/clean-architecture-telegram-bot/repositories"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongodb struct {
	db     *mongo.Database
	logger logger.Logger
}

// dsn returns formatted data source name.
func dsn(cfg configs.MongoDB) string {
	return fmt.Sprintf("%s://%s:%s@%s:%s",
		cfg.Protocol, cfg.Username, cfg.Password, cfg.Host, cfg.Port,
	)
}

// New creates a new connection to the database.
func New(ctx context.Context, cfg configs.MongoDB, logger logger.Logger) (repositories.DB, error) {
	clientOptions := options.Client().ApplyURI(dsn(cfg))

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to open the mongodb database driver: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to open the mongodb database driver: %v", err)
	}

	db := client.Database(cfg.DBName)

	return &mongodb{db: db, logger: logger}, nil
}

// Close terminates a database connection.
func (m *mongodb) Close(ctx context.Context) error {
	err := m.db.Client().Disconnect(ctx)
	if err != nil {
		return fmt.Errorf("can not close the database: %v", err)
	}

	return nil
}
