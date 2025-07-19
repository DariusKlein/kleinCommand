package subcommands

import (
	"context"
	"errors"
	"github.com/urfave/cli/v3"
	"log/slog"
)

var templateVar bool

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
		&cli.BoolFlag{
			Name:        "template",
			Aliases:     []string{"t"},
			Usage:       "usage",
			Value:       false,
			DefaultText: "errors template",
			Required:    false,
			Destination: &templateVar,
		},
	}
}

// templateAction logic for Template
func templateAction(context context.Context, c *cli.Command) error {
	slog.InfoContext(context, "template called", "logLevel", "INFO")
	slog.DebugContext(context, "template called", "logLevel", "DEBUG")
	if templateVar {
		slog.ErrorContext(context, "template called", "logLevel", "ERROR")
		return errors.New("template called with error")
	}
	return nil
}
