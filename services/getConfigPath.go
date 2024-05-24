package services

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

func GetConfigPath() (path string, configPath string, err error) {
	homeDir, _ := os.UserHomeDir()

	switch runtime.GOOS {
	case "windows":
		path = filepath.Dir(homeDir + "\\AppData\\Local\\kleinCommand\\")
	case "linux":
		path = filepath.Dir(homeDir + "/.config/kleinCommand")
	default:
		return "", "", errors.New("unsupported platform")
	}

	configPath = filepath.Join(path, "/config.toml")

	return path, configPath, nil
}
