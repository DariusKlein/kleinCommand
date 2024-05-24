package config

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/urfave/cli/v2"
	"os"
	"path/filepath"
	"runtime"
)

func subcommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:   "create",
			Usage:  "Generates a new configuration file",
			Action: creatAction,
		},
	}
}

func creatAction(c *cli.Context) error {
	var path string

	homeDir, _ := os.UserHomeDir()

	switch runtime.GOOS {
	case "windows":
		path = filepath.Dir(homeDir + "\\AppData\\Local\\kleinCommand\\")
	case "linux":
		path = filepath.Dir(homeDir + "/.config/kleinCommand")
	default:
		return errors.New("unsupported platform")
	}

	fmt.Println("Creating configuration file")
	if err := os.MkdirAll(path, 0770); err != nil {
		return err
	}
	client := github.NewClient(nil)
	_, _, _, _ = client.Repositories.GetContents(context.Background(), "DariusKlein", "kleinCommand", "config.toml", nil)
	configPath := filepath.Join(path, "/config.toml")

	err := os.WriteFile(configPath, []byte(""), 0644)
	if err != nil {
		return err
	}
	fmt.Println("Created: " + configPath)
	return nil
}
