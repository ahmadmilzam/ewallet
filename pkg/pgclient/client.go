package pgclient

import (
	"fmt"
	"log"
	"time"

	"github.com/ahmadmilzam/ewallet/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New() *sqlx.DB {
	dbConfig := config.GetDBConfig()

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbConfig.Username, dbConfig.Password, dbConfig.Name)
	db, err := sqlx.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("failure when opening db connection to: %s err: %v", dbConfig.GetConnectionURI(), err)
	}

	db.SetMaxIdleConns(dbConfig.Connection.MaxIdleConn)
	db.SetMaxOpenConns(dbConfig.Connection.MaxOpenConn)
	lifeTime := time.Second * time.Duration(dbConfig.Connection.MaxLifeTimeConn)
	db.SetConnMaxLifetime(lifeTime)

	return db
}
