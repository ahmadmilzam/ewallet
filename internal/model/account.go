package model

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
)

//go:generate mockery --name AccountSQLStore
type AccountSQLStore interface {
	Create(ctx context.Context, model *Account) (*Account, error)
	Update(ctx context.Context, model *Account) (*Account, error)
	DeleteByID(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (*Account, error)
	FindAll(ctx context.Context) ([]*Account, error)
}

type Account struct {
	ID        uuid.UUID `db:"id"`
	Phone     string    `db:"phone"`
	Role      string    `db:"role"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
