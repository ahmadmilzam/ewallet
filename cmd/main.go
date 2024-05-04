package main

import (
	"os"

	"github.com/ahmadmilzam/ewallet/config"
	"github.com/ahmadmilzam/ewallet/internal/console"
	"github.com/ahmadmilzam/ewallet/internal/migration"
	"github.com/ahmadmilzam/ewallet/internal/store"
	"github.com/ahmadmilzam/ewallet/internal/usecase"
	"github.com/ahmadmilzam/ewallet/pkg/logger"
	"github.com/urfave/cli/v2"
)

func main() {
	cliApp := &cli.App{}

	_ = config.Load("config", ".")
	appConfig := config.GetAppConfig()
	dbConfig := config.GetDBConfig()
	migrateCommand := migration.CreateMigrate(dbConfig.Name)

	pgstore := store.NewSQLStore()

	logger.InitializeLogger(logger.NewOption(logger.WithLevel(config.GetLogConfig().Level)))

	appUsecase := usecase.NewAppUsecase(pgstore, appConfig)

	cliApp.Commands = []*cli.Command{
		console.StartServer(appConfig, appUsecase),
		console.StartMigration(migrateCommand),
	}

	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}
}
