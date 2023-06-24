// Package keyboards containes all bot's keyboards.
package keyboards

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var MainKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Your Information", "information"),
		tgbotapi.NewInlineKeyboardButtonData("Help", "help"),
	),
)
