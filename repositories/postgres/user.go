// Package postgres implements postgres database functionalities.
package postgres

import (
	"context"
	"fmt"

	"github.com/barahouei/clean-architecture-telegram-bot/models"
)

// CreateUser creates a new user.
func (p *postgres) CreateUser(ctx context.Context, user *models.User) error {
	query := "INSERT INTO users(telegram_id, username, firstname, lastname, joined_at, language) VALUES($1, $2, $3, $4, $5, $6)"

	_, err := p.db.ExecContext(ctx, query, user.TelegramID, user.Username, user.FirstName, user.LastName, user.JoinedAt, user.Language)
	if err != nil {
		return err
	}

	return nil
}

// GetUser returns a user.
func (p *postgres) GetUser(ctx context.Context, telegramID int64) (*models.User, error) {
	user := models.User{}

	query := "SELECT * FROM users WHERE telegram_id=$1"

	row := p.db.QueryRowContext(ctx, query, telegramID)
	err := row.Scan(&user.ID, &user.TelegramID, &user.Username, &user.FirstName, &user.LastName, &user.JoinedAt, &user.Language)
	if err != nil {
		return nil, fmt.Errorf("can not fetch user %d from database: %v", telegramID, err)
	}

	return &user, err
}

// UpdateLanguage changes user's language.
func (p *postgres) UpdateLanguage(ctx context.Context, lang models.Language, telegramID int64) error {
	query := "UPDATE users SET language=$1 WHERE telegram_id=$2"

	_, err := p.db.ExecContext(ctx, query, lang, telegramID)
	if err != nil {
		return err
	}

	return nil
}
