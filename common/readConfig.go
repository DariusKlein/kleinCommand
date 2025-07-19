package common

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

var config Config

func ReadConfig() (Config, error) {
	_, configPath, err := GetConfigPath()
	if err != nil {
		fmt.Println(err)
	}

	file, err := os.ReadFile(configPath)
	if err != nil {
		return config, err
	}

	_, err = toml.Decode(string(file), &config)
	if err != nil {
		return config, err
	}

	return config, nil
}
