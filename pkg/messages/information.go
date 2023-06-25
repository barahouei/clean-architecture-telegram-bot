// Package messages contains all messages to be displayed.
package messages

import (
	"fmt"

	"github.com/barahouei/clean-architecture-telegram-bot/models"
)

// Information returns user's information message.
func Information(user *models.User) string {
	if user.Language == models.En {
		return fmt.Sprintf("Telegram ID: %d\nUsername: %s\nFirstname: %s\nLastname: %s\nLanguage: %v",
			user.TelegramID,
			user.Username,
			user.FirstName,
			user.LastName,
			user.Language.String(),
		)
	} else if user.Language == models.Fa {
		return fmt.Sprintf("آی‌دی تلگرام: %d\nنام کاربری: %s\nنام: %s\nنام خانوادگی: %s\nزبان: %v",
			user.TelegramID,
			user.Username,
			user.FirstName,
			user.LastName,
			user.Language.String(),
		)
	}

	return ""
}
