package subcommands

import (
	"context"
	"github.com/DariusKlein/kleinCommand/common"
	"github.com/urfave/cli/v3"
	"log/slog"
	"net"
)

// StopServiceTemplate Command
func StopServiceTemplate() *cli.Command {
	return &cli.Command{
		Name:      "stop service",
		Aliases:   []string{"stop"},
		Usage:     "template command usage",
		Action:    stopServiceTemplateAction,
		ArgsUsage: "args usage",
	}
}

// templateAction logic for Template
func stopServiceTemplateAction(context context.Context, c *cli.Command) error {
	conn, err := net.Dial("unix", common.ExampleServiceSocketPath)
	if err != nil {
		slog.ErrorContext(context, err.Error())
		return err
	}
	defer conn.Close()

	// Send the shutdown command.
	_, err = conn.Write([]byte("shutdown\n"))
	if err != nil {
		slog.ErrorContext(context, err.Error())
		return err
	}

	return nil
}
