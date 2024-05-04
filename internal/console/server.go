package console

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/ahmadmilzam/ewallet/config"
	restApi "github.com/ahmadmilzam/ewallet/internal/rest"
	"github.com/ahmadmilzam/ewallet/internal/usecase"
	"github.com/ahmadmilzam/ewallet/pkg/httpserver"
	"github.com/gorilla/mux"
	"github.com/urfave/cli/v2"
)

func StartServer(config config.AppConfig, usecase usecase.AppUsecaseInterface) *cli.Command {
	var err error

	return &cli.Command{
		Name:  "start",
		Usage: "Starting up application",
		Action: func(c *cli.Context) error {
			router := mux.NewRouter()
			restApi.RegisterRoutes(router, usecase)

			httpServer := httpserver.New(router, httpserver.WithPort(config.Port))
			httpServer.Start()

			slog.Info("Application started with runtime go version " + runtime.Version())
			// Waiting signal
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
	}
}
