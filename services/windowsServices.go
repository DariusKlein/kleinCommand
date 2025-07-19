//go:build windows

package services

//go:generate go build ./example

import _ "embed"

//go:embed exampleService.exe
var exampleService []byte
