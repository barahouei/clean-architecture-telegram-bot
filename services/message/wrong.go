// package account implements message request functionalities.
package message

import (
	"context"

	"github.com/barahouei/clean-architecture-telegram-bot/models"
	"github.com/barahouei/clean-architecture-telegram-bot/pkg/keyboards"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Wrong shows a wrong message.
func (m *msg) Wrong(ctx context.Context, msg tgbotapi.MessageConfig, lang models.Language) tgbotapi.MessageConfig {
	msg.Text = "Enter a valid message or command."
	msg.ReplyMarkup = keyboards.BackToMainKeyboard

	return msg
}
