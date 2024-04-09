package store

import (
	"log"
	"os"
	"testing"

	"github.com/ahmadmilzam/ewallet/config"
	"github.com/ahmadmilzam/ewallet/pkg/pgclient"
)

var testStore *Store

func TestMain(m *testing.M) {
	_ = config.Load("config", "../../config")
	sql := pgclient.New()

	if err := sql.DB.Ping(); err != nil {
		log.Fatal("cannot ping db: ", err)
	}

	testStore = &Store{
		DB:                sql,
		AccountQueryStore: NewAccountStore(sql),
	}

	os.Exit(m.Run())
}
