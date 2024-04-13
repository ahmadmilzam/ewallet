package store

import (
	"context"
	"database/sql"
)

type DBOperations interface {
	// Get(dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	// Select(dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	// Exec(query string, args ...any) (sql.Result, error)
	// ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}

func NewQueries(db DBOperations) *Queries {
	return &Queries{db: db}
}

type Queries struct {
	db DBOperations
}
