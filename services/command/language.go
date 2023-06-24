// package command implements command request functionalities.
package command

import (
	"context"

	"github.com/barahouei/clean-architecture-telegram-bot/pkg/keyboards"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Language shows a select language message.
func (m *cmd) Language(ctx context.Context, msg tgbotapi.MessageConfig) tgbotapi.MessageConfig {
	msg.Text = "یک زبان را انتخاب کنید:\n\nChoose a language:"
	msg.ReplyMarkup = keyboards.LangKeyboard

	return msg
}
