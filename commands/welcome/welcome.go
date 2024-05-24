package welcome

import (
	"fmt"
	"github.com/DariusKlein/kleinCommand/types"
	"github.com/urfave/cli/v2"
)

func Command(config types.Config) *cli.Command {

	return &cli.Command{
		Name:   "welcome",
		Usage:  "Explains cli tool",
		Action: action,
	}
}

func action(*cli.Context) error {
	fmt.Println("Welcome, \n you can use -h, --help for help.")
	return nil
}
