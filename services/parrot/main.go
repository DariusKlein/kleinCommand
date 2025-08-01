package main

import (
	"github.com/DariusKlein/kleinCommand/common"
	"github.com/DariusKlein/kleinCommand/services"
	"log"
	"net"
)

var socketPath = common.ParrotServiceSocketPath

func main() {
	services.BaseService(socketPath, func(command string, conn net.Conn) {
		_, err := conn.Write([]byte(command))
		if err != nil {
			log.Println(err.Error())
		}
	})
}
