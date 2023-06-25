// Package keyboards containes all bot's keyboards.
package keyboards

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	emoji "github.com/jayco/go-emoji-flag"
)

var LangKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(

		tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("English %v", emoji.GetFlag("US")), "en"),
		tgbotapi.NewInlineKeyboardButtonData(fmt.Sprintf("فارسی %v", emoji.GetFlag("IR")), "fa"),
	),
)
