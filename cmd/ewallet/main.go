package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ahmadmilzam/ewallet/config"
	"github.com/ahmadmilzam/ewallet/db"
	"github.com/ahmadmilzam/ewallet/handler"
	"github.com/ahmadmilzam/ewallet/pkg/logger"
	"github.com/ahmadmilzam/ewallet/pkg/statsd"
	"github.com/ahmadmilzam/ewallet/pkg/trace"
	"github.com/urfave/cli/v2"
)

const (
	ERR_NO_CHANGE = "no change"
)

func main() {

	cliApp := &cli.App{}

	_ = config.Load("config", "./config")

	var dbConfig = config.GetDBConfig()
	var migrate = db.CreateMigrate(dbConfig.Name)

	// logger.Init()
	logger.InitializeLogger(logger.NewOption(logger.WithLevel("debug")))
	statsd.Init()
	trace.Init()
	defer trace.Stop()

	cliApp.Commands = []*cli.Command{
		{
			Name:  "start",
			Usage: "Starting up ewallet",
			Action: func(c *cli.Context) error {
				ctx, cancel := context.WithCancel(context.Background())
				err := handler.Serve(ctx, config.GetServerAddress())
				if err != nil {
					panic(fmt.Sprintf("server exited with err: %s\n", err.Error()))
				}
				go callOnInterrupt(cancel)
				return nil
			},
		},
		{
			Name:        "migrate",
			Description: "migrate the database",
			Subcommands: []*cli.Command{
				{
					Name:        "create",
					Description: "create the migration file",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:     "filename",
							Usage:    "--filename create_user_table",
							Value:    "",
							Required: true,
						},
					},
					Action: func(c *cli.Context) error {
						if err := migrate.Create(c.String("filename")); err != nil {
							panic(fmt.Sprintf("Can't create db file with error: %v", err.Error()))
						}
						return nil
					},
				},
				{
					Name:        "up",
					Description: "run the migration files",
					Action: func(c *cli.Context) error {
						if err := migrate.Up(); err != nil && err.Error() != ERR_NO_CHANGE {
							panic(fmt.Sprintf("Can't run db up with error: %v", err.Error()))
						}
						return nil
					},
				},
				{
					Name:        "down",
					Description: "rollback the migration",
					Action: func(c *cli.Context) error {
						if err := migrate.Down(); err != nil && err.Error() != ERR_NO_CHANGE {
							panic(fmt.Sprintf("Can't run db down with error: %v", err.Error()))
						}
						return nil
					},
				},
			},
		},
	}

	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}
}

func callOnInterrupt(cancel context.CancelFunc) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	<-sigCh
	cancel()
}
