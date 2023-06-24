// Package keyboards containes all bot's keyboards.
package keyboards

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var BackToMainKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Back to main", "backToMain"),
	),
)
