package store

import (
	"fmt"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/pkg/pgclient"
	_ "github.com/lib/pq"
)

func NewStore() (*Store, error) {
	sql, err := pgclient.New()
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	if err := sql.DB.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return &Store{
		DB:                sql,
		AccountQueryStore: NewAccountStore(sql.DB),
	}, nil

	// alt version
	// return &Store{
	// 	AccountStore: &AccountStore{DB: db},
	// }, nil
}

type Store struct {
	DB                       *pgclient.Client
	entity.AccountQueryStore // TODO: add another store here and in model.Store interface
}

// alt version
// type Store struct {
// 	*AccountStore // TODO: add another store here and in model.Store interface
// }
