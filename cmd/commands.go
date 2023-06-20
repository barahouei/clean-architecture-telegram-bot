// Package cmd contains main functionality for the application using CLI commands.
package cmd

import (
	"fmt"
	"log"

	"github.com/urfave/cli/v2"
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

	return nil
}

// debug is a subcommand of the serve command that turns debug mode on.
func debug(c *cli.Context) error {
	log.Println("Debug mode is on")
	debugMode = true

	return nil
}
