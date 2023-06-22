// package account implements account functionalities.
package account

import (
	"context"
	"fmt"

	"github.com/barahouei/clean-architecture-telegram-bot/models"
)

// Create create a new user.
func (a *acc) Create(ctx context.Context, user *models.User) error {
	if a.IsExist(ctx, user) {
		return nil
	}

	err := a.db.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

// Get returns a user's full information.
func (a *acc) Get(ctx context.Context, user *models.User) (*models.User, error) {
	if !a.IsExist(ctx, user) {
		return nil, fmt.Errorf("user don't exist: %d", user.TelegramID)
	}

	user, err := a.db.GetUser(ctx, user.TelegramID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Language returns user's language.
func (a *acc) Language(ctx context.Context, user *models.User) (models.Language, error) {
	user, err := a.Get(ctx, user)
	if err != nil {
		return 0, err
	}

	return user.Language, nil
}

// IsExist checks if user is already a member.
func (a *acc) IsExist(ctx context.Context, user *models.User) bool {
	_, err := a.db.GetUser(ctx, user.TelegramID)

	return err == nil
}
