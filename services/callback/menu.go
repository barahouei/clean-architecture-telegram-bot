// package callback implements callback request functionalities.
package callback

import (
	"context"
	"fmt"

	"github.com/barahouei/clean-architecture-telegram-bot/models"
	"github.com/barahouei/clean-architecture-telegram-bot/pkg/keyboards"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Menu shows main menu.
func (c *call) Menu(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig {
	msg.Text = fmt.Sprintf("Welcome %s.", user.FirstName)
	msg.ReplyMarkup = keyboards.MainKeyboard

	return msg
}
