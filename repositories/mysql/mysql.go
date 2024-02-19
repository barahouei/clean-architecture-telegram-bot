// Package mysql implements mysql database functionalities.
package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/barahouei/clean-architecture-telegram-bot/configs"
	"github.com/barahouei/clean-architecture-telegram-bot/pkg/logger"
	"github.com/barahouei/clean-architecture-telegram-bot/repositories"

	_ "github.com/go-sql-driver/mysql"
)

type mysql struct {
	db     *sql.DB
	logger logger.Logger
}

// dsn returns formatted data source name.
func dsn(cfg configs.MySQL) string {
	return fmt.Sprintf("%s:%s@%s(%s:%s)/%s?parseTime=true",
		cfg.Username, cfg.Password, cfg.Protocol, cfg.Host, cfg.Port, cfg.DBName,
	)
}

// New creates a new connection to the database.
func New(ctx context.Context, cfg configs.MySQL, logger logger.Logger) (repositories.DB, error) {
	db, err := sql.Open("mysql", dsn(cfg))
	if err != nil {
		return nil, fmt.Errorf("failed to open the database driver: %v", err)
	}

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("can not ping the database: %v", err)
	}

	return &mysql{db: db, logger: logger}, nil
}

// Close terminates a database connection.
func (m *mysql) Close(ctx context.Context) error {
	err := m.db.Close()
	if err != nil {
		return fmt.Errorf("can not close the database: %v", err)
	}

	return nil
}
