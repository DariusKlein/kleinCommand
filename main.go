package main

import (
	"github.com/DariusKlein/kleinCommand/commands"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:      "KleinCommand",
		Usage:     "manage your home server",
		UsageText: "kleinCommand [category] [command] [arguments...]",
		Version:   "v0.0.1",
		Authors: []*cli.Author{
			{
				Name:  "Darius",
				Email: "darius.klein@dariusklein.nl",
			},
		},
		DefaultCommand: "help",
		Commands: []*cli.Command{
			commands.Template(),
			commands.Welcome(),
			commands.Boom(),
			commands.BubbleTeaTest(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
