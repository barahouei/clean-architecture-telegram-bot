// Package messages contains all messages to be displayed.
package messages

import (
	"fmt"

	"github.com/barahouei/clean-architecture-telegram-bot/models"
)

// Main returns main menu's message.
func Main(user *models.User) string {
	if user.Language == models.En {
		return fmt.Sprintf("Welcome %s.", user.FirstName)
	} else if user.Language == models.Fa {
		return fmt.Sprintf("خوش آمدید %s.", user.FirstName)
	}

	return ""
}
