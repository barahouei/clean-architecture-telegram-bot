package main

import (
	"log"

	"github.com/barahouei/clean-architecture-telegram-bot/cmd"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
