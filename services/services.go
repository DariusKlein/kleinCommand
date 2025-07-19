package services

import (
	"github.com/DariusKlein/kleinCommand/common"
	"os"
	"os/exec"
	"syscall"
)

func RunExampleService() error {
	return runService(common.ExampleServiceName, exampleService)
}

func runService(name string, file []byte) error {
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
