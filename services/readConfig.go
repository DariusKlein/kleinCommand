package services

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/DariusKlein/kleinCommand/types"
	"os"
)

var config types.Config

func ReadConfig() (types.Config, error) {
	_, configPath, err := GetConfigPath()
	if err != nil {
		fmt.Println(err)
	}

	file, err := os.ReadFile(configPath)
	if err != nil {
		return config, err
	}

	_, err = toml.Decode(string(file), &config)

	return config, nil
}
