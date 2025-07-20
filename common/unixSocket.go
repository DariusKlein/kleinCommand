package common

import (
	"net"
	"os"
	"path/filepath"
)

var ExampleServiceSocketPath = GetSocketPath(ExampleServiceName)
var ParrotServiceSocketPath = GetSocketPath(ParrotServiceName)

func GetSocketPath(serviceName string) string {
	return filepath.Join(os.TempDir(), serviceName+".sock")
}

func GetSocketConnection(socketPath string) (net.Conn, error) {
	return net.Dial("unix", socketPath)
}
