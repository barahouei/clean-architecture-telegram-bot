// Package postgres implements postgres database functionalities.
package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/barahouei/clean-architecture-telegram-bot/configs"
	"github.com/barahouei/clean-architecture-telegram-bot/pkg/logger"
	"github.com/barahouei/clean-architecture-telegram-bot/repositories"

	_ "github.com/lib/pq"
)

type postgres struct {
	db     *sql.DB
	logger logger.Logger
}

// dsn returns formatted data source name.
func dsn(cfg configs.Postgres) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSL,
	)
}

// New creates a new connection to the database.
func New(ctx context.Context, cfg configs.Postgres, logger logger.Logger) (repositories.DB, error) {
	db, err := sql.Open("postgres", dsn(cfg))
	if err != nil {
		return nil, fmt.Errorf("failed to open the database driver: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("can not ping the database: %v", err)
	}

	return &postgres{db: db, logger: logger}, nil
}

// Close terminates a database connection.
func (p *postgres) Close(ctx context.Context) error {
	err := p.db.Close()
	if err != nil {
		return fmt.Errorf("can not close the database: %v", err)
	}

	return nil
}
