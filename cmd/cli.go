// Package cmd contains main functionality for the application using CLI commands.
package cmd

import (
	"os"

	"github.com/urfave/cli/v2"
)

// Run executes the application.
func Run() error {
	app := cli.App{
		Name:      "Sample-telegram-bot",
		Usage:     "A sample project for creating telegram bots in Golang with clean architecture",
		UsageText: "Sample-telegram-bot [command] serve (s) [command option] debug (d)",
		Version:   "0.1",
		Commands: []*cli.Command{
			serveCMD,
		},
	}

	return app.Run(os.Args)
}
