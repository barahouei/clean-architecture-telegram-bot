// Package services implements all aplication services.
package services

import (
	"context"

	"github.com/barahouei/clean-architecture-telegram-bot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Command implements command functionalities.
type Command interface {
	Language(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig
}
