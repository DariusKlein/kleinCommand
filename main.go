package main

import (
	"context"
	config "github.com/DariusKlein/kleinCommand/commands/config"
	"github.com/DariusKlein/kleinCommand/commands/templateCommand"
	"github.com/DariusKlein/kleinCommand/common"
	"github.com/urfave/cli/v3"
	"log"
	"net/mail"
	"os"
)

var Config common.Config

func main() {
	Config, _ = common.ReadConfig()

	app := &cli.Command{
		Name:        "KleinCommand",
		Usage:       "CLI tool for internal use",
		UsageText:   "kleinCommand [category] [command] [arguments...]",
		Version:     "v0.1.0",
		HideVersion: true,
		Authors: []any{
			mail.Address{
				Name:    "Darius",
				Address: "darius.klein@dariusklein.nl",
			},
		},
		DefaultCommand: "help",
		Commands: []*cli.Command{
			config.Category(),
			templateCommand.Category(),
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
