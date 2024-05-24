package config

import (
	"github.com/DariusKlein/kleinCommand/types"
	"github.com/urfave/cli/v2"
)

func Command(config types.Config) *cli.Command {
	return &cli.Command{
		Name:            "config",
		Usage:           "manage config file",
		Aliases:         []string{"cf"},
		Action:          action,
		Category:        "\nManagement",
		Subcommands:     subcommands(),
		HideHelpCommand: true,
	}
}

func action(c *cli.Context) error {
	err := cli.ShowSubcommandHelp(c)
	if err != nil {
		return err
	}
	return nil
}
