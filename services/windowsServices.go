//go:build windows

package services

//go:generate go build -o binaries/ ./example
//go:generate go build -o binaries/ ./parrot

import (
	_ "embed"
	"os"
	"os/exec"
	"syscall"
)

//go:embed binaries/exampleService.exe
var exampleService []byte

//go:embed binaries/parrotService.exe
var parrotService []byte

func runService(name string, file []byte) error {
	executableName := name + ".exe"
	tempFile, err := os.CreateTemp("", executableName)
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
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	if err = cmd.Start(); err != nil {
		return err
	}

	return nil
}
