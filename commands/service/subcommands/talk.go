package subcommands

import (
	"bufio"
	"context"
	"github.com/DariusKlein/kleinCommand/common"
	"github.com/urfave/cli/v3"
	"log"
	"os"
)

// Talk Command
func Talk() *cli.Command {
	return &cli.Command{
		Name:      "talk",
		Usage:     "talk with parrot",
		Action:    templateAction,
		ArgsUsage: "send arg as message to parrot",
	}
}

// templateAction logic for Template
func templateAction(context context.Context, c *cli.Command) error {
	conn, err := common.GetSocketConnection(common.ParrotServiceSocketPath)
	if err != nil {
		return err
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	common.CatchInterrupt(func() {
		conn.Close()
		log.Fatal("Closed connection")
	})

	for scanner.Scan() {
		if _, err = conn.Write([]byte(scanner.Text() + "\n")); err != nil {
			return err
		}
		response := make([]byte, 1024)
		if _, err = conn.Read(response); err != nil {
			return err
		}
		log.Println("Response:", string(response))
	}
	return nil
}
