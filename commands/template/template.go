package template

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {

	return &cli.Command{
		Name:            "NAME",
		Usage:           "USAGE OF COMMAND",
		Aliases:         []string{"T"},
		Action:          action,
		Flags:           flags(),
		Category:        "\nTEMPLATE",
		Subcommands:     subcommands(),
		HideHelpCommand: true,
	}
}

func action(c *cli.Context) error {
	fmt.Println("TEMPLATE RESPONSE")
	return nil
}

func flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "test",
			Aliases: []string{"t"},
		},
	}
}

func subcommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:    "NAME",
			Usage:   "USAGE OF COMMAND",
			Aliases: []string{"T"},
			Action:  action,
			Flags:   []cli.Flag{},
		},
	}
}
