package common

import (
	"os"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func DeleteSelf() {
	path, _ := os.Executable()
	os.Remove(path)
}
