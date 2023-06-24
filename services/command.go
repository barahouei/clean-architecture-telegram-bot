// Package services implements all aplication services.
package services

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Command implements command functionalities.
type Command interface {
	Language(ctx context.Context, msg tgbotapi.MessageConfig) tgbotapi.MessageConfig
}
