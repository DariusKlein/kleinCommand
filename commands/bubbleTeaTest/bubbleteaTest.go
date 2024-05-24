package bubbleTeaTest

import (
	"github.com/DariusKlein/kleinCommand/types"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"
)

func Command(config types.Config) *cli.Command {
	return &cli.Command{
		Name:   "BubbleTeaTest",
		Usage:  "USAGE OF COMMAND",
		Action: action,
	}
}

func action(*cli.Context) error {
	p := tea.NewProgram(initialModel())
	p.Run()
	return nil
}

func initialModel() model {
	return model{
		// Our to-do list is a grocery list
		choices: []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}
