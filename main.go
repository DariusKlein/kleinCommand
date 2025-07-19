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
		log.Fatal(err)
	}

	setLogLevel(conf)

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
	opts := &slog.HandlerOptions{}
	switch strings.ToUpper(conf.Settings.LogLevel) {
	case "INFO":
		opts.Level = slog.LevelInfo
	case "WARN":
		opts.Level = slog.LevelWarn
	case "DEBUG":
		opts.Level = slog.LevelDebug
	case "ERROR":
		opts.Level = slog.LevelError
	}
	slog.SetLogLoggerLevel(opts.Level.Level())
}
