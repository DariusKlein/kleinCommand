package main

import (
	"github.com/DariusKlein/kleinCommand/commands/boom"
	"github.com/DariusKlein/kleinCommand/commands/bubbleTeaTest"
	"github.com/DariusKlein/kleinCommand/commands/config"
	"github.com/DariusKlein/kleinCommand/commands/template"
	"github.com/DariusKlein/kleinCommand/commands/welcome"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	config.CreatAction(&cli.Context{})
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
			template.Command(),
			welcome.Command(),
			boom.Command(),
			bubbleTeaTest.Command(),
			config.Command(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
