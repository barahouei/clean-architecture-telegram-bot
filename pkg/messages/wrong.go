// Package messages contains all messages to be displayed.
package messages

import (
	"github.com/barahouei/clean-architecture-telegram-bot/models"
)

// wrong returns wrong command or message.
func Wrong(user *models.User) string {
	if user.Language == models.En {
		return "Enter a valid message or command."
	} else if user.Language == models.Fa {
		return "یک پیام یا دستور معتبر وارد کنید."
	}

	return ""
}
