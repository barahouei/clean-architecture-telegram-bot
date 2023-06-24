// package callback implements callback request functionalities.
package callback

import (
	"github.com/barahouei/clean-architecture-telegram-bot/pkg/logger"
	"github.com/barahouei/clean-architecture-telegram-bot/repositories"
	"github.com/barahouei/clean-architecture-telegram-bot/services"
)

type call struct {
	db     repositories.DB
	logger logger.Logger
}

// New returns a callback service.
func New(db repositories.DB, logger logger.Logger) services.Callback {
	return &call{
		db:     db,
		logger: logger,
	}
}
