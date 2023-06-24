// package callback implements callback request functionalities.
package callback

import (
	"context"

	"github.com/barahouei/clean-architecture-telegram-bot/models"
	"github.com/barahouei/clean-architecture-telegram-bot/pkg/keyboards"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Help shows help message.
func (c *call) Help(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig {
	msg.Text = "This is a clean architecture telegram bot in Golang!"
	msg.ReplyMarkup = keyboards.BackToMainKeyboard

	return msg
}
