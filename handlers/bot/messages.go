// Package bot handles bot activities.
package bot

import (
	"context"

	"github.com/barahouei/clean-architecture-telegram-bot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Message handles message requests.
func (h *handler) Message(ctx context.Context, tgbot *tgbotapi.BotAPI, update tgbotapi.Update) {
	deleteLastMessage(tgbot, update.Message.Chat.ID, update.Message.MessageID)

	user := &models.User{
		TelegramID: update.Message.From.ID,
		Username:   update.Message.From.UserName,
		FirstName:  update.Message.From.FirstName,
		LastName:   update.Message.From.LastName,
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	var err error
	user.Language, err = h.account.Language(ctx, user)
	if err != nil {
		h.logger.Error(err)
	}

	msg = h.message.Wrong(ctx, msg, user.Language)

	_, err = tgbot.Send(msg)
	if err != nil {
		h.logger.Error(err)
	}
}
