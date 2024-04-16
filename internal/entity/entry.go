package entity

import (
	"context"
	"time"
)

//go:generate mockery --name TransferStoreQuerier
type EntryQuery interface {
	CreateTransfer(ctx context.Context, model Transfer) (Transfer, error)
	UpdateTransfer(ctx context.Context, model Transfer) (Transfer, error)
	DeleteTransfer(ctx context.Context, id string) error
	FindTransfer(ctx context.Context, id string) (Transfer, error)
	FindTransfers(ctx context.Context) ([]Transfer, error)
}

type Entry struct {
	ID            string    `db:"id"`
	WalletID      string    `db:"wallet_id"`
	CreditAmount  int64     `db:"credit_amount"`
	DebitAmount   int64     `db:"debit_amount"`
	BalanceBefore float64   `db:"balance_before"`
	BalanceAfter  float64   `db:"balance_after"`
	CorrelationID string    `db:"correlation_id"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}
