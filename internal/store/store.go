package store

import (
	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/pkg/pgclient"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//go:generate mockery --name "SQLStoreInterface" --output "./_mock" --outpkg "mockery"
type SQLStoreInterface interface {
	entity.AccountQuery
	entity.WalletQuery
	entity.TransferCounterQuery
	entity.TransferQuery
	entity.EntryQuery
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	*sqlx.DB // for DB Transaction purposes
	*QueryCommands
}

// NewStore creates a new store
func NewSQLStore() SQLStoreInterface {
	sql := pgclient.New()
	return &SQLStore{
		DB:            sql,
		QueryCommands: NewQueryCommands(sql),
	}
}
