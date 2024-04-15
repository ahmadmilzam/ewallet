package entity

import (
	"context"
	"time"
)

type TransferCounterQuery interface {
	CreateCounter(ctx context.Context, tc *TransferCounter) (*TransferCounter, error)
	UpdateCounter(ctx context.Context, tc *TransferCounter) (*TransferCounter, error)
	FindCounterById(ctx context.Context, id string) (*TransferCounter, error)
	FindCounterForUpdateById(ctx context.Context, id string) (*TransferCounter, error)
}

type TransferCounter struct {
	WalletId            string    `db:"wallet_id"`
	CountDaily          int16     `db:"count_daily"`
	CountMonthly        int16     `db:"count_monthly"`
	CreditAmountDaily   float64   `db:"amount_daily"`
	CreditAmountMonthly float64   `db:"amount_monthly"`
	CreatedAt           time.Time `db:"created_at"`
	UpdatedAt           time.Time `db:"updated_at"`
}
