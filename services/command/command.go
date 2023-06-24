// package command implements command request functionalities.
package command

import (
	"github.com/barahouei/clean-architecture-telegram-bot/pkg/logger"
	"github.com/barahouei/clean-architecture-telegram-bot/repositories"
	"github.com/barahouei/clean-architecture-telegram-bot/services"
)

type cmd struct {
	db     repositories.DB
	logger logger.Logger
}

// New returns a command service.
func New(db repositories.DB, logger logger.Logger) services.Command {
	return &cmd{
		db:     db,
		logger: logger,
	}
}
