package store

import (
	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/pkg/pgclient"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type SQLStoreInterface interface {
	entity.StoreQuerier
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	*sqlx.DB
	*Queries
}

// NewStore creates a new store
func NewSQLStore() SQLStoreInterface {
	sql := pgclient.New()
	return &SQLStore{
		DB:      sql,
		Queries: NewQueries(sql),
	}
}
