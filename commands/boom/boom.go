package boom

import (
	"fmt"
	"github.com/DariusKlein/kleinCommand/types"
	"github.com/urfave/cli/v2"
)

func Command(config types.Config) *cli.Command {

	return &cli.Command{
		Name:    "boom",
		Usage:   "explode",
		Aliases: []string{"b"},
		Action:  action,
	}
}

func action(c *cli.Context) error {
	fmt.Println("BOOM")
	return nil
}
