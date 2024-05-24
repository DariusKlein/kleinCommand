package games

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/DariusKlein/kleinCommand/services"
	"github.com/DariusKlein/kleinCommand/types"
	"github.com/urfave/cli/v2"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func subcommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:   "register",
			Usage:  "Register game server",
			Action: registerAction,
			Flags:  registerFlags(),
		},
		{
			Name:    "list",
			Usage:   "List of game servers",
			Aliases: []string{"l"},
			Action:  listAction,
			Flags:   []cli.Flag{},
		},
		{
			Name:    "status",
			Usage:   "Get status of game servers",
			Aliases: []string{"s"},
			Action:  statusAction,
			Flags:   []cli.Flag{},
		},
	}
}

func registerAction(c *cli.Context) error {

	for _, games := range config.CustomCommands.Games {
		if games.Name == c.String("name") {
			return fmt.Errorf("game '%s' already exists", games.Name)
		}
	}

	config.CustomCommands.Games = append(config.CustomCommands.Games, types.Games{
		Name:          c.String("name"),
		StatusCommand: c.String("status"),
		StartCommand:  c.String("start_command"),
		StopCommand:   c.String("stop_command"),
	})

	configString, err := toml.Marshal(config)
	if err != nil {
		return err
	}

	_, configPath, err := services.GetConfigPath()
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile(configPath, configString, 0644)
	if err != nil {
		return err
	}
	return nil
}

func registerFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Usage:    "Game server name",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "status_command",
			Aliases:  []string{"status"},
			Usage:    "Game server status command",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "start_command",
			Aliases:  []string{"start"},
			Usage:    "Game server start command",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "stop_command",
			Aliases:  []string{"stop"},
			Usage:    "Game server stop command",
			Required: true,
		},
	}
}

func listAction(c *cli.Context) error {
	fmt.Println("List of game servers")
	for i, game := range config.CustomCommands.Games {
		fmt.Println(
			strconv.Itoa(i+1) + ") \n" +
				"\tName: " + game.Name +
				"\n\tStatus command: " + game.StatusCommand +
				"\n\tStart command: " + game.StartCommand +
				"\n\tStop command: " + game.StopCommand)
	}
	return nil
}

func statusAction(c *cli.Context) error {
	fmt.Println("List of game servers")
	for _, game := range config.CustomCommands.Games {
		command := strings.Fields(game.StatusCommand)
		cmd := exec.Command(command[0], command[1:]...)
		stdout, err := cmd.Output()
		if err != nil {
			return err
		}
		fmt.Println(game.Name + ": " + string(stdout))
	}

	return nil
}
