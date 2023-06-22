// Package cmd contains main functionality for the application using CLI commands.
package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/barahouei/clean-architecture-telegram-bot/configs"
	"github.com/barahouei/clean-architecture-telegram-bot/models"
	"github.com/barahouei/clean-architecture-telegram-bot/pkg/logger/zap"
	"github.com/barahouei/clean-architecture-telegram-bot/repositories/postgres"
	"github.com/barahouei/clean-architecture-telegram-bot/services/account"
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
	fmt.Printf("Debug mod is: %t\n", debugMode)
	fmt.Println("clean-architecture-telegram-bot")

	file, err := os.OpenFile("logs/bot.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	logger := zap.New(file, zapcore.InfoLevel)

	logger.Info("Bot Starting...")

	cfg, err := configs.New(logger, models.Development)
	if err != nil {
		return err
	}

	fmt.Println(cfg)

	db, err := postgres.New(cfg.Postgres, logger)
	if err != nil {
		logger.Error(err)
		return err
	}
	defer db.Close()

	user := &models.User{
		TelegramID: 650,
		Username:   "test5",
		FirstName:  "name",
		LastName:   "family",
		JoinedAt:   time.Now(),
		Language:   models.En,
	}
	// err = db.CreateUser(context.Background(), user)
	// if err != nil {
	// 	return err
	// }

	// user, err := db.GetUser(context.Background(), 649)
	// if err != nil {
	// 	return err
	// }

	acc := account.New(db, logger)

	// acc.Create(context.Background(), user)

	user, err = acc.Get(context.Background(), user)
	if err != nil {
		logger.Error(err)

		return err
	}
	fmt.Println(acc.Language(context.Background(), user))
	fmt.Println(acc.IsExist(context.Background(), user))
	fmt.Println(user)

	return nil
}

// debug is a subcommand of the serve command that turns debug mode on.
func debug(c *cli.Context) error {
	log.Println("Debug mode is on")
	debugMode = true

	return nil
}
