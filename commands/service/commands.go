package service

import (
	"context"
	"github.com/DariusKlein/kleinCommand/commands/service/subcommands"
	"github.com/urfave/cli/v3"
)

// Category service
func Category() *cli.Command {
	return &cli.Command{
		Name:            "service",
		Usage:           "commands for starting, stopping and communicating with services",
		Action:          Action,
		Commands:        commands(),
		HideHelpCommand: true,
	}
}

// commands for service Category
func commands() []*cli.Command {
	return []*cli.Command{
		parrot(),
	}
}

// Action show help command if no sub commands are given for Category
func Action(context context.Context, c *cli.Command) error {
	return cli.ShowSubcommandHelp(c)
}

// parrot sub-category of service Category
func parrot() *cli.Command {
	return &cli.Command{
		Name:   "parrot",
		Usage:  "commands for interacting with parrot service",
		Action: Action,
		Commands: []*cli.Command{
			subcommands.StartParrotService(),
			subcommands.StopParrotService(),
			subcommands.Talk(),
		},
		HideHelpCommand: true,
	}
}
