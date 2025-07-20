package subcommands

import (
	"context"
	"github.com/DariusKlein/kleinCommand/common"
	"github.com/urfave/cli/v3"
	"log/slog"
)

// StopParrotService Command
func StopParrotService() *cli.Command {
	return &cli.Command{
		Name:   "stop",
		Usage:  "stop parrot service",
		Action: stopParrotServiceAction,
	}
}

// templateAction logic for Template
func stopParrotServiceAction(context context.Context, c *cli.Command) error {
	conn, err := common.GetSocketConnection(common.ParrotServiceSocketPath)
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
