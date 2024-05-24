package config

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/DariusKlein/kleinCommand/services"
	"github.com/google/go-github/github"
	"github.com/urfave/cli/v2"
	"os"
)

func subcommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:   "create",
			Usage:  "Generates a new configuration file",
			Action: creatAction,
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "force",
					Aliases: []string{"f"},
					Usage:   "force overwrite",
				},
			},
		},
		{
			Name:   "get",
			Usage:  "prints configuration file to stdout",
			Action: printAction,
		},
	}
}

func creatAction(c *cli.Context) error {
	path, configPath, err := services.GetConfigPath()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Creating configuration file")
	if err = os.MkdirAll(path, 0770); err != nil {
		return err
	}

	if _, err = os.Stat(configPath); errors.Is(err, os.ErrNotExist) || c.Bool("force") {

		client := github.NewClient(nil)
		defaultConfig, _, _, _ := client.Repositories.GetContents(context.Background(), "DariusKlein", "kleinCommand", "config.toml", nil)

		configString, _ := base64.StdEncoding.DecodeString(*defaultConfig.Content)

		err = os.WriteFile(configPath, configString, 0644)
		if err != nil {
			return err
		}
	}
	fmt.Println("Created: " + configPath)
	return nil
}

func printAction(c *cli.Context) error {
	_, configPath, err := services.GetConfigPath()
	if err != nil {
		fmt.Println(err)
	}

	file, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	fmt.Println(string(file))
	return nil
}
