package main

import (
	"bufio"
	"github.com/DariusKlein/kleinCommand/common"
	"log"
	"net"
	"os"
)

var socketPath = common.ExampleServiceSocketPath

func main() {
	common.CatchInterrupt(func() {
		os.Remove(socketPath)
	})

	if common.FileExists(socketPath) {
		log.Fatal("Socket file exists.")
	}

	// Create the UNIX socket listener.
	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatalf("Failed to listen on socket: %v", err)
	}

	defer listener.Close()

	log.Println("Service started, listening on", socketPath)

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error: %v", err)
			return
		}
		defer conn.Close()

		reader := bufio.NewReader(conn)
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Read error: %v", err)
			return
		}

		if command == "shutdown\n" {
			log.Println("Shutdown command received, exiting.")
			// This will cause the listener.Accept() to unblock and the program to exit.
			listener.Close()
		}
	}()

	for {
		if _, err := os.Stat(socketPath); os.IsNotExist(err) {
			break
		}
	}
	common.DeleteSelf()
}
