package cmd

import (
	"fmt"

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

func serve(c *cli.Context) error {
	fmt.Printf("Debug mod is: %t\n", debugMode)
	fmt.Println("clean-architecture-telegram-bot")

	return nil
}

func debug(c *cli.Context) error {
	fmt.Println("Debug mode is on")
	debugMode = true

	return nil
}
