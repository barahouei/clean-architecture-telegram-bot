// package callback implements callback request functionalities.
package callback

import (
	"context"

	"github.com/barahouei/clean-architecture-telegram-bot/models"
	"github.com/barahouei/clean-architecture-telegram-bot/pkg/keyboards"
	"github.com/barahouei/clean-architecture-telegram-bot/pkg/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Language shows a select language menu.
func (m *call) Language(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig {
	msg.Text = messages.Language(user)
	msg.ReplyMarkup = keyboards.Language()

	return msg
}
