// Package messages contains all messages to be displayed.
package messages

import (
	"github.com/barahouei/clean-architecture-telegram-bot/models"
)

// language returns a choosing language message.
func Language(user *models.User) string {
	return "یک زبان را انتخاب کنید:\n\nChoose a language:"
}
