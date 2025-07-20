package main

import (
	"context"
	"github.com/DariusKlein/kleinCommand/commands/config"
	"github.com/DariusKlein/kleinCommand/commands/templateCommand"
	"github.com/DariusKlein/kleinCommand/common"
	"github.com/urfave/cli/v3"
	"log"
	"log/slog"
	"net/mail"
	"os"
	"strings"
)

func main() {
	conf, err := common.ReadConfig()
	if err != nil {
		slog.Error(err.Error())
	} else {
		setLogLevel(conf)
	}

	app := &cli.Command{
		Name:        "KleinCommand",
		Usage:       "CLI tool for internal use",
		UsageText:   "kleinCommand [category] [command] [arguments...]",
		Version:     "v0.1.0",
		HideVersion: true,
		Authors: []any{
			mail.Address{
				Name:    "Darius",
				Address: "darius.klein@dariusklein.nl",
			},
		},
		DefaultCommand: "help",
		Commands: []*cli.Command{
			config.Category(),
			templateCommand.Category(),
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func setLogLevel(conf common.Config) {
	var logLevel slog.Level
	switch strings.ToUpper(conf.Settings.LogLevel) {
	case "INFO":
		logLevel = slog.LevelInfo
	case "WARN":
		logLevel = slog.LevelWarn
	case "DEBUG":
		logLevel = slog.LevelDebug
	case "ERROR":
		logLevel = slog.LevelError
	default:
		log.Fatal("unknown log level", logLevel, conf.Settings.LogLevel)
	}
	slog.SetLogLoggerLevel(logLevel)
}
