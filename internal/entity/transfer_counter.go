package entity

import (
	"context"
	"time"
)

type TransferCounterQuery interface {
	CreateCounter(ctx context.Context, counter *TransferCounter) (*TransferCounter, error)
	UpdateCounter(ctx context.Context, counter *TransferCounter) (*TransferCounter, error)
	FindCounterById(ctx context.Context, id string) (*TransferCounter, error)
	FindCounterForUpdateById(ctx context.Context, id string) (*TransferCounter, error)
}

type TransferCounter struct {
	WalletId            string    `db:"wallet_id"`
	CreditCountDaily    int16     `db:"credit_count_daily"`
	CreditCountMonthly  int16     `db:"credit_count_monthly"`
	CreditAmountDaily   float64   `db:"credit_amount_daily"`
	CreditAmountMonthly float64   `db:"credit_amount_monthly"`
	CreatedAt           time.Time `db:"created_at"`
	UpdatedAt           time.Time `db:"updated_at"`
}
