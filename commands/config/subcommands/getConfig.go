package subcommands

import (
	"context"
	"fmt"
	"github.com/DariusKlein/kleinCommand/common"
	"github.com/urfave/cli/v3"
	"os"
)

// GetConfig Command
func GetConfig() *cli.Command {
	return &cli.Command{
		Name:   "get",
		Usage:  "read configuration file",
		Action: getConfigAction,
		Flags:  getConfigFlags(),
	}
}

// getConfigFlags Register cli flags
func getConfigFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "force",
			Aliases:     []string{"f"},
			Usage:       "force overwrite",
			Destination: &force,
		},
	}
}

// getConfigAction logic for GetConfig
func getConfigAction(context context.Context, c *cli.Command) error {
	_, configPath, err := common.GetConfigPath()
	if err != nil {
		return err
	}

	file, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	fmt.Println(string(file))
	return nil
}
