package main

import (
	"github.com/DariusKlein/kleinCommand/common"
	"github.com/DariusKlein/kleinCommand/services"
)

var socketPath = common.ExampleServiceSocketPath

func main() {
	services.BaseService(socketPath, func(command string) {
		switch command {
		case "example\n":

		}
	})
}
