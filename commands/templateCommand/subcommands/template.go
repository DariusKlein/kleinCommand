package subcommands

import (
	"context"
	"github.com/urfave/cli/v3"
)

var templateVar int

// Template Command
func Template() *cli.Command {
	return &cli.Command{
		Name:      "template command",
		Usage:     "template command usage",
		Action:    templateAction,
		Flags:     templateFlags(),
		ArgsUsage: "args usage",
	}
}

// templateFlags Register cli flags
func templateFlags() []cli.Flag {
	return []cli.Flag{
		&cli.IntFlag{
			Name:        "template",
			Aliases:     []string{"t"},
			Usage:       "usage",
			Value:       1000,
			DefaultText: "default text",
			Required:    false,
			Destination: &templateVar,
		},
	}
}

// templateAction logic for Template
func templateAction(context context.Context, c *cli.Command) error {
	return nil
}
