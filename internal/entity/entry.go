package entity

import (
	"context"
	"time"
)

type EntryQuery interface {
	CreateEntry(ctx context.Context, model *Entry) (*Entry, error)
	FindEntryById(ctx context.Context, id string) (*Entry, error)
}

type Entry struct {
	ID            string    `db:"id"`
	WalletID      string    `db:"wallet_id"`
	CreditAmount  int64     `db:"credit_amount"`
	DebitAmount   int64     `db:"debit_amount"`
	BalanceBefore int64     `db:"balance_before"`
	BalanceAfter  int64     `db:"balance_after"`
	CorrelationID string    `db:"correlation_id"`
	TransferID    string    `db:"transfer_id"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}
