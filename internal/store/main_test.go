package store_test

import (
	"log"
	"os"
	"testing"

	"github.com/ahmadmilzam/ewallet/config"
	"github.com/ahmadmilzam/ewallet/internal/store"
	"github.com/ahmadmilzam/ewallet/pkg/pgclient"
)

var testStore *store.SQLStore

func TestMain(m *testing.M) {
	_ = config.Load("config", "../../")
	sql := pgclient.New()

	if err := sql.DB.Ping(); err != nil {
		log.Fatal("cannot ping db: ", err)
	}

	testStore = &store.SQLStore{
		DB:      sql,
		Queries: store.NewQueries(sql),
	}

	os.Exit(m.Run())
}
