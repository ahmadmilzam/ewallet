package store

import (
	"context"
	"database/sql"
)

type QueryCommands struct {
	db SQLOperations
}

type SQLOperations interface {
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}

func NewQueryCommands(db SQLOperations) *QueryCommands {
	return &QueryCommands{db: db}
}
