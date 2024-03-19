package main

import (
	"os"

	"github.com/ahmadmilzam/ewallet/internal/config"
	"github.com/ahmadmilzam/ewallet/internal/console"
	"github.com/ahmadmilzam/ewallet/pkg/logger"
	"github.com/ahmadmilzam/ewallet/pkg/sqlclient"
	"github.com/ahmadmilzam/ewallet/pkg/statsd"
	"github.com/ahmadmilzam/ewallet/pkg/trace"
	"github.com/urfave/cli"
)

func main() {
	cliApp := cli.NewApp()

	_ = config.Load("config", "./configs")
	logger.Init()
	statsd.Init()
	trace.Init()
	defer trace.Stop()

	sqlClient := sqlclient.New()
	defer sqlClient.Close()

	cliApp.Commands = []cli.Command{
		console.StartServer(),
		console.Migration(sqlClient.DB.DB),
	}
	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}
}
