// Package services implements all aplication services.
package services

import (
	"context"

	"github.com/barahouei/clean-architecture-telegram-bot/models"
)

// Account implements user account functionalities.
type Account interface {
	Create(ctx context.Context, user *models.User) error
	Get(ctx context.Context, user *models.User) (*models.User, error)
	Language(ctx context.Context, user *models.User) (models.Language, error)
	ChooseLanguage(ctx context.Context, lang models.Language, user *models.User) error
	IsExist(ctx context.Context, user *models.User) bool
}
