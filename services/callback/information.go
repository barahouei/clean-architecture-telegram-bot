// package callback implements callback request functionalities.
package callback

import (
	"context"

	"github.com/barahouei/clean-architecture-telegram-bot/models"
	"github.com/barahouei/clean-architecture-telegram-bot/pkg/keyboards"
	"github.com/barahouei/clean-architecture-telegram-bot/pkg/messages"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Information shows user's information.
func (c *call) Information(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig {
	msg.Text = messages.Information(user)
	msg.ReplyMarkup = keyboards.BackToMain(user.Language)

	return msg
}
