package main

import (
	"github.com/DariusKlein/kleinCommand/common"
	"github.com/DariusKlein/kleinCommand/services"
	"net"
)

var socketPath = common.ExampleServiceSocketPath

func main() {
	services.BaseService(socketPath, func(command string, conn net.Conn) {})
}
