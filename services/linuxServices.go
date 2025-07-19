//go:build linux

package services

//go:generate go build ./example

import _ "embed"

//go:embed exampleService
var exampleService []byte
