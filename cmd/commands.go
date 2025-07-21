// Package cmd contains main functionality for the application using CLI commands.
package cmd

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/barahouei/clean-architecture-telegram-bot/configs"
	"github.com/barahouei/clean-architecture-telegram-bot/handlers/bot"
	"github.com/barahouei/clean-architecture-telegram-bot/models"
	"github.com/barahouei/clean-architecture-telegram-bot/pkg/logger/zap"
	"github.com/barahouei/clean-architecture-telegram-bot/repositories/database"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap/zapcore"
)

var (
	serveCMD = &cli.Command{
		Name:    "serve",
		Aliases: []string{"s"},
		Usage:   "Serving telegram bot",
		Action:  serve,
		Subcommands: []*cli.Command{
			debugCMD,
		},
	}

	debugCMD = &cli.Command{
		Name:    "debug",
		Aliases: []string{"d"},
		Usage:   "Turn on debug mode",
		After:   serve,
		Action:  debug,
	}

	debugMode = false
)

// serve is the main command that runs the application.
func serve(c *cli.Context) error {
	createLogFolderErr := os.Mkdir("logs", os.ModePerm)
	if createLogFolderErr != nil {
		if !os.IsExist(createLogFolderErr) {
			slog.Error("failed to create logs directory", slog.Any("error", createLogFolderErr))

			os.Exit(1)
		}
	}

	file, createLogFileErr := os.OpenFile("logs/bot.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if createLogFileErr != nil {
		slog.Error("failed to create log file", slog.Any("error", createLogFileErr))

		os.Exit(1)
	}
	defer file.Close()

	logger := zap.New(file, zapcore.InfoLevel)

	cfg, cfgErr := configs.New(logger, models.Development)
	if cfgErr != nil {
		slog.Error("failed to load config", slog.Any("error", cfgErr))

		os.Exit(1)
	}

	ctx := context.TODO()

	db, dbErr := database.NewDatabase(ctx, cfg, logger)
	if dbErr != nil {
		logger.Error(fmt.Sprintf("failed to initialize database: %v", dbErr))

		os.Exit(1)
	}

	defer db.Close(ctx)

	bot := bot.New(db, logger)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		startErr := bot.Start(ctx, cfg.App, debugMode)
		if startErr != nil {
			logger.Error(fmt.Sprintf("bot failed to start: %v", startErr))

			os.Exit(1)
		}
	}()

	<-sigChan

	logger.Info("Received an interrupt or terminate signal, Bot stopped...")

	return nil
}

// debug is a subcommand of the serve command that turns debug mode on.
func debug(c *cli.Context) error {
	log.Println("Debug mode is on")

	debugMode = true

	return nil
}
