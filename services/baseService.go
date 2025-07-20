package services

import (
	"bufio"
	"github.com/DariusKlein/kleinCommand/common"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

func BaseService(socketPath string, logic func(string)) {
	// Remove socket path on os.Interrupt and syscall.SIGTERM
	common.CatchInterrupt(func() {
		os.Remove(socketPath)
	})
	// Check if socket exists
	if common.FileExists(socketPath) {
		log.Fatal("Socket file exists.")
	}
	// Create the UNIX socket listener.
	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatalf("Failed to listen on socket: %v", err)
	}

	log.Println("Service started, listening on", socketPath)

	shutdownChan := make(chan struct{})

	var once sync.Once
	closeOnce := func() {
		once.Do(func() {
			close(shutdownChan)
		})
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Printf("Accept error: %v", err)
				closeOnce()
				return
			}
			go handleConnection(conn, logic, listener)
		}
	}()

	go func() {
		ticker := time.NewTicker(5 * time.Second) // Check every 5 seconds
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				if _, err := os.Stat(socketPath); os.IsNotExist(err) {
					log.Println("Socket file deleted externally, signaling shutdown.")
					closeOnce()
					return
				}
			case <-shutdownChan:
				return
			}
		}
	}()

	<-shutdownChan

	common.DeleteSelf()
}

func handleConnection(conn net.Conn, logic func(string), listener net.Listener) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "shutdown" {
			log.Println("Shutdown command received, exiting.")
			listener.Close()
		} else {
			logic(input)
		}
	}
}
