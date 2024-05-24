package main

import (
	"github.com/urfave/cli/v2"
	"kleinCommand/commands"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:           "KleinCommand",
		Usage:          "manage your home server",
		UsageText:      "kleinCommand [category] [command] [arguments...]",
		DefaultCommand: "help",
		Commands: []*cli.Command{
			commands.Welcome(),
			commands.Boom(),
			commands.BubbleTeaTest(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
