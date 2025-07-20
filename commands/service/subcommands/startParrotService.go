package subcommands

import (
	"context"
	"github.com/DariusKlein/kleinCommand/services"
	"github.com/urfave/cli/v3"
)

// StartParrotService Command
func StartParrotService() *cli.Command {
	return &cli.Command{
		Name:   "start",
		Usage:  "start parrot service",
		Action: startParrotServiceAction,
	}
}

// startServiceTemplateAction logic for StartServiceTemplate
func startParrotServiceAction(context context.Context, c *cli.Command) error {
	return services.RunParrotService()
}
