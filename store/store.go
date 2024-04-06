package store

import (
	"fmt"

	"github.com/ahmadmilzam/ewallet/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewStore(datasource string) (*Store, error) {
	db, err := sqlx.Open("postgres", datasource)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return &Store{
		AccountStore: NewAccountStore(db),
	}, nil

	// alt version
	// return &Store{
	// 	AccountStore: &AccountStore{DB: db},
	// }, nil
}

type Store struct {
	model.AccountStore // TODO: add another store here and in model.Store interface
}

// alt version
// type Store struct {
// 	*AccountStore // TODO: add another store here and in model.Store interface
// }
