package common

type Config struct {
	Settings Settings `toml:"settings"`
}

type Settings struct {
	Environment string `toml:"environment"`
	LogLevel    string `toml:"log_level"`
}
