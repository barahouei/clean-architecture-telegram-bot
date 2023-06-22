// Package repositories implements database functionalities.
package repositories

import (
	"context"

	"github.com/barahouei/clean-architecture-telegram-bot/models"
)

// DB implements all database functionalities.
type DB interface {
	User

	Close() error
}

// User implements user functionalities.
type User interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, telegramID int64) (*models.User, error)
}
