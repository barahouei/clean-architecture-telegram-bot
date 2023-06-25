// Package messages contains all messages to be displayed.
package messages

import (
	"github.com/barahouei/clean-architecture-telegram-bot/models"
)

// Help returns a help message.
func Help(user *models.User) string {
	if user.Language == models.En {
		return "This is a clean architecture telegram bot in Golang!"
	} else if user.Language == models.Fa {
		return "این یک بات تلگرام معماری تمیز است در گولنگ!"
	}

	return ""
}
