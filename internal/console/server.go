package console

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ahmadmilzam/ewallet/config"
	"github.com/ahmadmilzam/ewallet/internal/api"
	"github.com/ahmadmilzam/ewallet/internal/api/httpserver"
	"github.com/ahmadmilzam/ewallet/internal/api/middleware"
	"github.com/ahmadmilzam/ewallet/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

func StartServer(config config.AppConfig, usecase usecase.AppUsecaseInterface) *cli.Command {
	var err error

	return &cli.Command{
		Name:  "start",
		Usage: "Starting up ewallet",
		Action: func(c *cli.Context) error {
			handler := gin.New()
			handler.Use(middleware.RequestLog())
			handler.Use(gin.Recovery())
			api.NewRouter(handler, usecase)

			httpServer := httpserver.New(handler, httpserver.WithPort(config.Port))
			httpServer.Start()

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
