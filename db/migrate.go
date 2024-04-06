package db

import (
	"fmt"
	"os"
	"time"

	"github.com/ahmadmilzam/ewallet/db/sqlclient"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Migrations interface {
	Up() error
	Down() error
	Create(title string) error
}

func CreateMigrate(databaseName string) Migrations {
	return &PostgresMigrations{
		sourceFile:   "db/migrations/",
		databaseName: databaseName,
	}
}

type PostgresMigrations struct {
	migrate      *migrate.Migrate
	sourceFile   string
	databaseName string
}

func (p *PostgresMigrations) init() error {
	if p.migrate != nil {
		fmt.Println("p.migrate not nil")
		return nil
	}

	sql := sqlclient.New()
	defer sql.Close()

	sourceFile := fmt.Sprintf("file://%s", p.sourceFile)
	driver, err := postgres.WithInstance(sql.DB.DB, &postgres.Config{})

	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(sourceFile, p.databaseName, driver)
	if err != nil {
		return err
	}

	p.migrate = m

	return nil
}

func (p *PostgresMigrations) Up() error {
	if err := p.init(); err != nil {
		fmt.Println("error init PG migrate")
		return err
	}

	return p.migrate.Up()
}

func (p *PostgresMigrations) Down() error {
	if err := p.init(); err != nil {
		return err
	}

	return p.migrate.Steps(-1)
}

func (p *PostgresMigrations) Create(title string) error {
	if title == "" {
		return errors.New("Title can't be empty")
	}
	fileNameUp, fileNameDown := p.generateFileName(title)

	if _, err := os.Create(fileNameUp); err != nil {
		return err
	}

	if _, err := os.Create(fileNameDown); err != nil {
		_ = os.Remove(fileNameUp)
		return err
	}

	return nil
}

func (p PostgresMigrations) generateFileName(title string) (fileNameUp string, fileNameDown string) {
	now := time.Now()
	unixTime := now.Unix()

	fileNameUp = fmt.Sprintf("%s/%d_%s.up.sql", p.sourceFile, unixTime, title)
	fileNameDown = fmt.Sprintf("%s/%d_%s.down.sql", p.sourceFile, unixTime, title)

	return fileNameUp, fileNameDown
}
