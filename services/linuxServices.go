//go:build linux

package services

//go:generate go build -o binaries/ ./example
//go:generate go build -o binaries/ ./parrot

import (
	_ "embed"
	"errors"
	"github.com/DariusKlein/kleinCommand/common"
	"os"
	"os/exec"
	"syscall"
)

//go:embed binaries/exampleService
var exampleService []byte

//go:embed binaries/parrotService
var parrotService []byte

func runService(name string, file []byte) error {
	// check for existing socket
	if common.FileExists(common.GetSocketPath(name)) {
		return errors.New("File " + common.GetSocketPath(name) + " already exists.")
	}
	tempFile, err := os.CreateTemp("", name)
	if err != nil {
		return err
	}

	if _, err = tempFile.Write(file); err != nil {
		tempFile.Close()
		return err
	}
	if err = tempFile.Close(); err != nil {
		return err
	}
	if err = os.Chmod(tempFile.Name(), 0777); err != nil {
		return err
	}

	cmd := exec.Command(tempFile.Name())
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
	if err = cmd.Start(); err != nil {
		return err
	}

	return nil
}
