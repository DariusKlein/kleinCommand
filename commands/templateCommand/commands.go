package templateCommand

import (
	"context"
	"github.com/DariusKlein/kleinCommand/commands/templateCommand/subcommands"
	"github.com/urfave/cli/v3"
)

// Category CATEGORY NAME
func Category() *cli.Command {
	return &cli.Command{
		Name:            "template",
		Usage:           "template commands",
		Aliases:         []string{"template"},
		Action:          Action,
		Commands:        commands(),
		HideHelpCommand: true,
	}
}

// commands for CATEGORY NAME Category
func commands() []*cli.Command {
	return []*cli.Command{
		//commands or sub-category here
		subCategory(),
	}
}

// action show help command if no sub commands are given for Category
func Action(context context.Context, c *cli.Command) error {
	err := cli.ShowSubcommandHelp(c)
	if err != nil {
		return err
	}
	return nil
}

// PLACEHOLDER sub-category of CATEGORY NAME Category
func subCategory() *cli.Command {
	return &cli.Command{
		Name:   "sub-category-template",
		Usage:  "commands for sub-category-template",
		Action: Action,
		Commands: []*cli.Command{
			//commands or sub-category here
			subcommands.Template(),
		},
		HideHelpCommand: true,
	}
}
