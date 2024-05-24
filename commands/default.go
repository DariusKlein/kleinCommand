package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func Welcome() *cli.Command {

	return &cli.Command{
		Name:   "welcome",
		Usage:  "Explains cli tool",
		Action: welcome,
	}
}

func welcome(*cli.Context) error {
	fmt.Println("Welcome, \n you can use -h, --help for help.")
	return nil
}
