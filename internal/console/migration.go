package console

import (
	"database/sql"
	"fmt"

	"github.com/ahmadmilzam/ewallet/internal/config"
	"github.com/ahmadmilzam/ewallet/internal/db"
	"github.com/urfave/cli"
)

const (
	ERR_NO_CHANGE = "no change"
)

func Migration(d *sql.DB) cli.Command {
	var config = config.GetDBConfig()
	var migrate = db.CreateMigrate(d, "postgres", config.Name)

	return cli.Command{
		Name:        "migrate",
		Description: "migrate the database",
		Subcommands: []cli.Command{
			{
				Name:        "create",
				Description: "create the migration file",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:     "filename",
						Usage:    "--filename create_user_table",
						Value:    "",
						Required: true,
					},
				},
				Action: func(c *cli.Context) {
					if err := migrate.Create(c.String("filename")); err != nil {
						panic(fmt.Sprintf("Can't create db file with error: %v", err.Error()))
					}
				},
			}, {
				Name:        "up",
				Description: "run the migration files",
				Action: func(c *cli.Context) {
					if err := migrate.Up(); err != nil && err.Error() != ERR_NO_CHANGE {
						panic(fmt.Sprintf("Can't run db up with error: %v", err.Error()))
					}
				},
			}, {
				Name:        "down",
				Description: "rollback the migration",
				Action: func(c *cli.Context) {
					if err := migrate.Down(); err != nil && err.Error() != ERR_NO_CHANGE {
						panic(fmt.Sprintf("Can't run db down with error: %v", err.Error()))
					}
				},
			},
		},
	}
}
