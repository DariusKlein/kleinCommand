package subcommands

import (
	"context"
	"github.com/DariusKlein/kleinCommand/services"
	"github.com/urfave/cli/v3"
)

// StartServiceTemplate Command
func StartServiceTemplate() *cli.Command {
	return &cli.Command{
		Name:      "start service",
		Aliases:   []string{"start"},
		Usage:     "template command usage",
		Action:    startServiceTemplateAction,
		ArgsUsage: "args usage",
	}
}

// startServiceTemplateAction logic for StartServiceTemplate
func startServiceTemplateAction(context context.Context, c *cli.Command) error {
	return services.RunExampleService()
}
