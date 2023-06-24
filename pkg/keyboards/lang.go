// Package keyboards containes all bot's keyboards.
package keyboards

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var LangKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("English", "en"),
		tgbotapi.NewInlineKeyboardButtonData("فارسی", "fa"),
	),
)
