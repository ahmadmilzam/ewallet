package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ahmadmilzam/ewallet/config"
	"github.com/ahmadmilzam/ewallet/internal/api"
	"github.com/ahmadmilzam/ewallet/internal/api/httpserver"
	"github.com/ahmadmilzam/ewallet/internal/migration"
	"github.com/ahmadmilzam/ewallet/internal/store"
	"github.com/ahmadmilzam/ewallet/internal/usecase"
	"github.com/ahmadmilzam/ewallet/pkg/logger"
	"github.com/ahmadmilzam/ewallet/pkg/statsd"
	"github.com/ahmadmilzam/ewallet/pkg/trace"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

const (
	ERR_NO_CHANGE = "no change"
)

func main() {

	cliApp := &cli.App{}

	_ = config.Load("config", "./config")

	appConfig := config.GetAppConfig()
	dbConfig := config.GetDBConfig()
	migrate := migration.CreateMigrate(dbConfig.Name)

	pgstore, err := store.NewStore()

	if err != nil {
		panic(err)
	}

	// logger.Init()
	logger.InitializeLogger(logger.NewOption(logger.WithLevel("debug")))
	statsd.Init()
	trace.Init()
	defer trace.Stop()

	accountUsecase := usecase.NewAccountUsecase(pgstore)

	// Passing also the basic auth middleware to all  Routers -.

	cliApp.Commands = []*cli.Command{
		{
			Name:  "start",
			Usage: "Starting up ewallet",
			Action: func(c *cli.Context) error {
				handler := gin.Default()
				api.NewRouter(handler, accountUsecase)

				httpServer := httpserver.New(handler, httpserver.Port(appConfig.Port))
				httpServer.Start()
				// Waiting signal -.
				interrupt := make(chan os.Signal, 1)
				signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

				select {
				case s := <-interrupt:
					slog.Info("app run", "signal", s.String())
				case err = <-httpServer.Notify():
					slog.Error("app interupted", fmt.Errorf("httpServer.Notify: %w", err))
				}

				// Shutdown
				err = httpServer.Shutdown()

				if err != nil {
					slog.Error("app shutdown", fmt.Errorf("httpServer.Shutdown: %w", err))
				}
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
