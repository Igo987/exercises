package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"

	"github.com/Igo87/project/config"
	"github.com/Igo87/project/logger"
	masker "github.com/Igo87/project/masker"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "masker",
		Usage: "masker",
		Authors: []*cli.Author{
			{
				Name:  "Igor",
				Email: "guruGOlang@gmail.com",
			},
		},
	}
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "path_read",
			Value: "./src/links.txt",
			Usage: "path to read file",
		},
		&cli.StringFlag{
			Name:  "path_write",
			Value: "./src/links.txt",
			Usage: "path to write file",
		},
		&cli.StringFlag{
			Name:  "log_level",
			Value: "slog.LevelDebug",
			Usage: "logLevel",
		},
	}
	app.Action = func(c *cli.Context) error {
		err := doAllWorks(c)
		if err != nil {
			return err
		}
		return nil
	}

	app.Commands = []*cli.Command{
		{
			Name:  "run",
			Usage: "run the service",
			Action: func(c *cli.Context) error {
				err := doAllWorks(c)
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		logger.LogStart("slog.LevelError").Error("error when running the app", err)
	}

}

func doAllWorks(c *cli.Context) error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	go masker.Stop(ctx)
	log := logger.LogStart(c.String("log_level"))
	cfg, err := config.ReadConfig("./config/config.yaml")
	if err != nil {
		log.Error("error when reading the config file", err)
	}

	logFile, err := openLogFile(cfg.GetPathToLogFile())
	if err != nil {
		log.Error("error when opening the log file", err)
	}
	defer logFile.Close()

	logger := logger.WriteLogInTheFile(logFile)
	logger.LogAttrs(ctx, slog.LevelDebug, "programm was started")

	somePres := masker.NewPresent()
	somePres.Path = c.String("path_read")
	someProd := masker.NewProduce()
	someProd.Path = c.String("path_write")
	newService := masker.NewService(somePres, someProd)

	if err := newService.Run(); err != nil {
		logger.Error("error when running the service", err)
	} else {
		logger.Info("successfully finished")
		log.Info("successfully finished")
	}
	return nil
}

func openLogFile(filePath string) (*os.File, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}
	return file, nil
}
