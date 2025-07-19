package config

import (
	"context"
	"github.com/DariusKlein/kleinCommand/commands/config/subcommands"
	"github.com/urfave/cli/v3"
)

// Category config
func Category() *cli.Command {
	return &cli.Command{
		Name:            "config",
		Usage:           "command to interact with config",
		Action:          Action,
		Commands:        commands(),
		HideHelpCommand: true,
	}
}

// commands for Config Category
func commands() []*cli.Command {
	return []*cli.Command{
		subcommands.CreateConfig(),
		subcommands.GetConfig(),
	}
}

// Action show help command if no sub commands are given for Config
func Action(context context.Context, c *cli.Command) error {
	return cli.ShowSubcommandHelp(c)
}
