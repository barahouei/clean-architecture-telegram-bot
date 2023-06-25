// package account implements message request functionalities.
package message

import (
	"context"

	"github.com/barahouei/clean-architecture-telegram-bot/models"
	"github.com/barahouei/clean-architecture-telegram-bot/pkg/keyboards"
	"github.com/barahouei/clean-architecture-telegram-bot/pkg/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Wrong shows a wrong message.
func (m *msg) Wrong(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig {
	msg.Text = messages.Wrong(user)
	msg.ReplyMarkup = keyboards.BackToMain(user.Language)

	return msg
}
