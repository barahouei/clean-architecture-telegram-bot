// package account implements account functionalities.
package account

import (
	"github.com/barahouei/clean-architecture-telegram-bot/pkg/logger"
	"github.com/barahouei/clean-architecture-telegram-bot/repositories"
	"github.com/barahouei/clean-architecture-telegram-bot/services"
)

type acc struct {
	db     repositories.DB
	logger logger.Logger
}

// New returns a new user account service.
func New(db repositories.DB, logger logger.Logger) services.Account {
	return &acc{
		db:     db,
		logger: logger,
	}
}
