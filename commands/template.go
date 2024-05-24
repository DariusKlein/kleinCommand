package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func Template() *cli.Command {

	return &cli.Command{
		Name:   "NAME",
		Usage:  "USAGE OF COMMAND",
		Action: template,
	}
}

func template(*cli.Context) error {
	fmt.Println("TEMPLATE RESPONSE")
	return nil
}
