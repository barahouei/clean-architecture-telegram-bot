// Package keyboards containes all bot's keyboards.
package keyboards

import (
	"github.com/barahouei/clean-architecture-telegram-bot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// BackToMain returns back to main keyboard.
func BackToMain(lang models.Language) tgbotapi.InlineKeyboardMarkup {
	if lang == models.En {
		return backToMainEn
	} else if lang == models.Fa {
		return backToMainFa
	}

	return tgbotapi.InlineKeyboardMarkup{}
}

// backToMainEn is back to main keyboard in English.
var backToMainEn = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Back to main", "backToMain"),
	),
)

// backToMainFa is back to main keyboard in Persian.
var backToMainFa = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("برگشت به منوی اصلی", "backToMain"),
	),
)
