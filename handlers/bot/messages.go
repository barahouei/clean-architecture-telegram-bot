// Package bot handles bot activities.
package bot

import (
	"context"

	"github.com/barahouei/clean-architecture-telegram-bot/pkg/keyboards"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Message handles message requests.
func (h *handler) Message(ctx context.Context, tgbot *tgbotapi.BotAPI, update tgbotapi.Update) {
	deleteLastMessage(tgbot, update.Message.Chat.ID, update.Message.MessageID)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	msg.Text = "Enter a valid command."
	msg.ReplyMarkup = keyboards.BackToMainKeyboard

	_, err := tgbot.Send(msg)
	if err != nil {
		h.logger.Error(err)
	}

}
