// Package services implements all aplication services.
package services

import (
	"context"

	"github.com/barahouei/clean-architecture-telegram-bot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Callback implements callback functionalities.
type Callback interface {
	Menu(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig
	Help(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig
	Information(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig
	Language(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig
}
