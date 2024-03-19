package console

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ahmadmilzam/ewallet/internal/config"
	"github.com/ahmadmilzam/ewallet/internal/handler"
	"github.com/urfave/cli"
)

func StartServer() cli.Command {
	return cli.Command{
		Name:        "start",
		Description: "Starting up ewallet",
		Action: func(c *cli.Context) {
			ctx, cancel := context.WithCancel(context.Background())
			if err := handler.Serve(ctx, config.GetServerAddress(), handler.Dependencies{}); err != nil {
				panic(fmt.Sprintf("server exited with err: %s\n", err.Error()))
			}
			go callOnInterrupt(cancel)
		},
	}
}

func callOnInterrupt(cancel context.CancelFunc) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	<-sigCh
	cancel()
}
