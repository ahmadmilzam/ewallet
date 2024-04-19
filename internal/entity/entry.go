package entity

import (
	"context"
	"time"
)

//go:generate mockery --name TransferStoreQuerier
type EntryQuery interface {
	CreateEntry(ctx context.Context, model *Entry) (*Entry, error)
}

type Entry struct {
	ID            string    `db:"id"`
	WalletID      string    `db:"wallet_id"`
	CreditAmount  float64   `db:"credit_amount"`
	DebitAmount   float64   `db:"debit_amount"`
	BalanceBefore float64   `db:"balance_before"`
	BalanceAfter  float64   `db:"balance_after"`
	CorrelationID string    `db:"correlation_id"`
	TransferID    string    `db:"transfer_id"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}
