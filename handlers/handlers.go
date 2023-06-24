// package handlers handles all incoming requests.
package handlers

import (
	"context"

	"github.com/barahouei/clean-architecture-telegram-bot/configs"
)

// Bot implements bot functionalities.
type Bot interface {
	Start(ctx context.Context, cfg configs.App, debugMode bool) error
}
