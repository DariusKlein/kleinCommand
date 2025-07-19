package common

import (
	"os"
	"os/signal"
	"syscall"
)

func CatchInterrupt(logic func()) {
	cmd := make(chan os.Signal, 1)
	signal.Notify(cmd, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-cmd
		logic()
		os.Exit(0)
	}()
}
