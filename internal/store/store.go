package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/pkg/pgclient"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type contextProp string

const (
	Tx contextProp = "tx"
)

func NewStore() (*Store, error) {
	sql := pgclient.New()

	// defer sql.DB.Close()

	if err := sql.DB.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %w", err)
	}

	return &Store{
		DB:                sql,
		AccountQueryStore: NewAccountStore(sql),
	}, nil

	// alt version
	// return &Store{
	// 	AccountStore: &AccountStore{DB: db},
	// }, nil
}

type Store struct {
	DB                       *sqlx.DB
	entity.AccountQueryStore // TODO: add another store here and in model.Store interface
}

func (s *Store) BeginTx(ctx context.Context) context.Context {
	tx, _ := s.DB.DB.Begin()
	ctx = context.WithValue(ctx, Tx, tx)
	return ctx
}

func (s *Store) CommitTx(ctx context.Context) error {
	tx, ok := ctx.Value(Tx).(*sql.Tx)
	if !ok {
		return errors.New("failed to commit on non transaction mode")
	}

	return tx.Commit()
}

func (s *Store) RollbackTx(ctx context.Context) error {
	tx, ok := ctx.Value(Tx).(*sql.Tx)
	if !ok {
		return errors.New("failed to rollback on non transaction mode")
	}
	tx.Rollback()
	return nil
}

// alt version
// type Store struct {
// 	*AccountStore // TODO: add another store here and in model.Store interface
// }
