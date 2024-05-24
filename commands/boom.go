package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func Boom() *cli.Command {

	return &cli.Command{
		Name:    "boom",
		Usage:   "explode",
		Aliases: []string{"b"},
		Action:  boom,
	}
}

func boom(c *cli.Context) error {
	fmt.Println("BOOM")
	return nil
}
