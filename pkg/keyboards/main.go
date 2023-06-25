// Package keyboards containes all bot's keyboards.
package keyboards

import (
	"github.com/barahouei/clean-architecture-telegram-bot/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Main returns maun menu's keyboard.
func Main(lang models.Language) tgbotapi.InlineKeyboardMarkup {
	if lang == models.En {
		return mainEN
	} else if lang == models.Fa {
		return mainFa
	}

	return tgbotapi.InlineKeyboardMarkup{}
}

// mainEN is main menu in English
var mainEN = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Your Information", "information"),
		tgbotapi.NewInlineKeyboardButtonData("Help", "help"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Choosing another language", "language"),
	),
)

// mainFa is main menu in Persian.
var mainFa = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("اطلاعات شما", "information"),
		tgbotapi.NewInlineKeyboardButtonData("کمک", "help"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("انتخاب زبان دیگر", "language"),
	),
)
