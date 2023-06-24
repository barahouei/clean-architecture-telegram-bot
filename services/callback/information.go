// package callback implements callback request functionalities.
package callback

import (
	"context"
	"fmt"

	"github.com/barahouei/clean-architecture-telegram-bot/models"
	"github.com/barahouei/clean-architecture-telegram-bot/pkg/keyboards"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Information shows user's information.
func (c *call) Information(ctx context.Context, msg tgbotapi.MessageConfig, user *models.User) tgbotapi.MessageConfig {
	info := fmt.Sprintf("Telegram ID: %d\nUsername: %s\nFirstname: %s\nLastname: %s\nLanguage:%v",
		user.TelegramID,
		user.Username,
		user.FirstName,
		user.LastName,
		user.Language.String(),
	)

	msg.Text = info
	msg.ReplyMarkup = keyboards.BackToMainKeyboard

	return msg
}
