// Package bot handles bot activities.
package bot

import (
	"context"
	"fmt"

	"github.com/barahouei/clean-architecture-telegram-bot/models"
	"github.com/barahouei/clean-architecture-telegram-bot/pkg/keyboards"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Callback handles callback requests.
func (h *handler) Callback(ctx context.Context, tgbot *tgbotapi.BotAPI, update tgbotapi.Update) {
	deleteLastMessage(tgbot, update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)

	user := &models.User{
		TelegramID: update.CallbackQuery.From.ID,
		Username:   update.CallbackQuery.From.UserName,
		FirstName:  update.CallbackQuery.From.FirstName,
		LastName:   update.CallbackQuery.From.LastName,
	}

	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)

	switch update.CallbackQuery.Data {
	case "en":
		err := h.account.ChooseLanguage(ctx, models.En, user)
		if err != nil {
			h.logger.Error(err)
		}

		msg.Text = mainMenu(user.FirstName)
		msg.ReplyMarkup = keyboards.MainKeyboard
	case "fa":
		err := h.account.ChooseLanguage(ctx, models.Fa, user)
		if err != nil {
			h.logger.Error(err)
		}

		msg.Text = mainMenu(user.FirstName)
		msg.ReplyMarkup = keyboards.MainKeyboard
	case "information":
		lang, err := h.account.Language(ctx, user)
		if err != nil {
			h.logger.Error(err)
		}

		info := fmt.Sprintf("Telegram ID: %d\nUsername: %s\nFirstname: %s\nLastname: %s\nLanguage:%v",
			user.TelegramID,
			user.Username,
			user.FirstName,
			user.LastName,
			lang.String(),
		)

		msg.Text = info
		msg.ReplyMarkup = keyboards.BackToMainKeyboard
	case "help":
		msg.Text = "This is a telegram bot!"
		msg.ReplyMarkup = keyboards.BackToMainKeyboard

	case "backToMain":
		msg.Text = mainMenu(user.FirstName)
		msg.ReplyMarkup = keyboards.MainKeyboard
	}

	_, err := tgbot.Send(msg)
	if err != nil {
		h.logger.Error(err)
	}
}
