// Package cmd contains main functionality for the application using CLI commands.
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/barahouei/clean-architecture-telegram-bot/configs"
	"github.com/barahouei/clean-architecture-telegram-bot/pkg/logger/zap"
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

const (
	_ configs.Mode = iota
	development
	production
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

	cfg, err := configs.New(logger, development)
	if err != nil {
		return err
	}

	fmt.Println(cfg)

	return nil
}

// debug is a subcommand of the serve command that turns debug mode on.
func debug(c *cli.Context) error {
	log.Println("Debug mode is on")
	debugMode = true

	return nil
}
