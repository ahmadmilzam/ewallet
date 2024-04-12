package store

import (
	"context"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/pkg/pgclient"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type TransferTxParams struct {
	FromAccountID int64 `json:"src_account_id"`
	ToAccountID   int64 `json:"dst_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Journal     entity.Journal  `json:"journal"`
	FromAccount entity.Account  `json:"src_account"`
	ToAccount   entity.Account  `json:"dst_account"`
	SrcTransfer entity.Transfer `json:"src_transfer"`
	DstTransfer entity.Transfer `json:"dst_transfer"`
}

type Store interface {
	entity.StoreQuerier
	CreateAccountTx(ctx context.Context, a entity.Account, w entity.Wallet) error
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	*sqlx.DB
	*Queries
}

// NewStore creates a new store
func NewStore() Store {
	sql := pgclient.New()

	return &SQLStore{
		DB:      sql,
		Queries: NewQueries(sql),
	}
	// alt version
	// return &Store{
	// 	AccountStore: &AccountStore{DB: db},
	// }, nil
}

// alt version
// type Store struct {
// 	*AccountStore // TODO: add another store here and in model.Store interface
// }
