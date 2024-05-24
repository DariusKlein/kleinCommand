package main

import (
	"github.com/DariusKlein/kleinCommand/commands/boom"
	"github.com/DariusKlein/kleinCommand/commands/bubbleTeaTest"
	"github.com/DariusKlein/kleinCommand/commands/config"
	"github.com/DariusKlein/kleinCommand/commands/games"
	"github.com/DariusKlein/kleinCommand/commands/template"
	"github.com/DariusKlein/kleinCommand/commands/welcome"
	"github.com/DariusKlein/kleinCommand/services"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	Config, _ := services.ReadConfig()

	app := &cli.App{
		Name:        "KleinCommand",
		Usage:       "manage your home server",
		UsageText:   "kleinCommand [category] [command] [arguments...]",
		Version:     "v0.0.1",
		HideVersion: true,
		Authors: []*cli.Author{
			{
				Name:  "Darius",
				Email: "darius.klein@dariusklein.nl",
			},
		},
		DefaultCommand: "help",
		Commands: []*cli.Command{
			template.Command(Config),
			welcome.Command(Config),
			boom.Command(Config),
			bubbleTeaTest.Command(Config),
			config.Command(Config),
			games.Command(Config),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
