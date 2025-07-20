package services

import (
	"github.com/DariusKlein/kleinCommand/common"
)

func RunExampleService() error {
	return runService(common.ExampleServiceName, exampleService)
}
