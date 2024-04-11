package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

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
		DB:           sql,
		Tx:           nil,
		AccountQuery: NewAccountStore(sql),
	}, nil

	// alt version
	// return &Store{
	// 	AccountStore: &AccountStore{DB: db},
	// }, nil
}

type Store struct {
	*sqlx.DB
	*sqlx.Tx
	entity.AccountQuery
	entity.WalletQuery
	entity.JournalQuery
	entity.TransferQuery // TODO: add another store here and in model.Store interface
}

func (s *Store) BeginTx(ctx context.Context) error {
	tx, err := s.DB.BeginTxx(ctx, nil)
	if err != nil {
		//log errors with context
		slog.Error("fail to begin tx", "error", err)
		return err
	}
	s.Tx = tx
	return nil
}

func (s *Store) CommitTx(ctx context.Context) error {
	err := s.Tx.Commit()
	if err != nil {
		//log errors with context
		slog.Error("fail to commit tx", "error", err)
	}
	return err
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
