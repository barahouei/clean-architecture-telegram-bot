package main

import (
	"log/slog"
	"os"

	"github.com/barahouei/clean-architecture-telegram-bot/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		slog.Error("failed to run command", slog.Any("error", err))

		os.Exit(1)
	}
}
