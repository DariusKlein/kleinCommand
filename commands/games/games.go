package games

import (
	"github.com/DariusKlein/kleinCommand/types"
	"github.com/urfave/cli/v2"
)

var config types.Config

func Command(conf types.Config) *cli.Command {
	config = conf
	return &cli.Command{
		Name:            "games",
		Usage:           "manage game servers",
		Aliases:         []string{"gs"},
		Action:          action,
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
