package common

import (
	"os"
	"path/filepath"
)

var ExampleServiceSocketPath = GetSocketPath(ExampleServiceName)

func GetSocketPath(serviceName string) string {
	return filepath.Join(os.TempDir(), serviceName+".sock")
}
