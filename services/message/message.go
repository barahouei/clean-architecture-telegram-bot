// package account implements message request functionalities.
package message

import (
	"github.com/barahouei/clean-architecture-telegram-bot/pkg/logger"
	"github.com/barahouei/clean-architecture-telegram-bot/repositories"
	"github.com/barahouei/clean-architecture-telegram-bot/services"
)

type msg struct {
	db     repositories.DB
	logger logger.Logger
}

// New returns a message service.
func New(db repositories.DB, logger logger.Logger) services.Message {
	return &msg{
		db:     db,
		logger: logger,
	}
}
