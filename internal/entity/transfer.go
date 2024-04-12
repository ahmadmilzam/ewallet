package entity

import (
	"context"
	"time"
)

//go:generate mockery --name TransferStoreQuerier
type TransferQuery interface {
	CreateTransfer(ctx context.Context, model Transfer) (Transfer, error)
	UpdateTransfer(ctx context.Context, model Transfer) (Transfer, error)
	DeleteTransfer(ctx context.Context, id string) error
	FindTransfer(ctx context.Context, id string) (Transfer, error)
	FindTransfers(ctx context.Context) ([]Transfer, error)
}

type Transfer struct {
	ID            string    `json:"id" db:"id"`
	WalletID      string    `json:"wallet_id" db:"wallet_id"`
	CreditAmount  int64     `json:"credit_amount" db:"credit_amount"`
	DebitAmount   int64     `json:"debit_amount" db:"debit_amount"`
	CorrelationID string    `json:"correlation" db:"correlation_id"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}
