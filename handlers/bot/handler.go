// Package bot handles bot activities.
package bot

import (
	"fmt"

	"github.com/barahouei/clean-architecture-telegram-bot/pkg/logger"
	"github.com/barahouei/clean-architecture-telegram-bot/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// handler handles all services.
type handler struct {
	logger  logger.Logger
	account services.Account
	message services.Message
	command services.Command
}

// deleteLastMessage deletes recieved request for user.
func deleteLastMessage(tgbot *tgbotapi.BotAPI, chatID int64, messageID int) {
	deleteLastMessage := tgbotapi.NewDeleteMessage(chatID, messageID)

	tgbot.Request(deleteLastMessage)
}

// mainMenu is main menu's message.
func mainMenu(firstname string) string {
	return fmt.Sprintf("Welcome %s", firstname)
}
