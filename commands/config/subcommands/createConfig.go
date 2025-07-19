package subcommands

import (
	"context"
	"errors"
	"fmt"
	"github.com/DariusKlein/kleinCommand/common"
	"github.com/urfave/cli/v3"
	"log"
	"os"
)

var force bool

// CreateConfig Command
func CreateConfig() *cli.Command {
	return &cli.Command{
		Name:   "create",
		Usage:  "Generates a new configuration file",
		Action: createConfigAction,
		Flags:  createConfigFlags(),
	}
}

// createConfigFlags Register cli flags
func createConfigFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "force",
			Aliases:     []string{"f"},
			Usage:       "force overwrite",
			Destination: &force,
		},
	}
}

// createConfigAction create config file
func createConfigAction(context context.Context, c *cli.Command) error {
	path, configPath, err := common.GetConfigPath()
	if err != nil {
		return err
	}
	log.Println("Creating configuration file")
	if err = os.MkdirAll(path, 0770); err != nil {
		return err
	}

	if _, err = os.Stat(configPath); errors.Is(err, os.ErrNotExist) || force {

		err = os.WriteFile(configPath, common.DefaultConfig, 0644)
		if err != nil {
			return err
		}
		fmt.Println("Created: " + configPath)
	} else {
		fmt.Println("Configuration file already exists")
	}
	return nil
}
