package entity

import (
	"context"
	"time"
)

type TransferCounterQuery interface {
	CreateCounter(ctx context.Context, counter *TransferCounter) (*TransferCounter, error)
	UpdateCounter(ctx context.Context, counter *TransferCounter) error
	FindCounterById(ctx context.Context, id string) (*TransferCounter, error)
	FindCounterForUpdateById(ctx context.Context, id string) (*TransferCounter, error)
}

type TransferCounter struct {
	WalletID            string    `db:"wallet_id"`
	CreditCountDaily    int16     `db:"credit_count_daily"`
	CreditCountMonthly  int16     `db:"credit_count_monthly"`
	CreditAmountDaily   int64     `db:"credit_amount_daily"`
	CreditAmountMonthly int64     `db:"credit_amount_monthly"`
	CreatedAt           time.Time `db:"created_at"`
	UpdatedAt           time.Time `db:"updated_at"`
}

type UpdateTransferCounter struct {
	WalletID      string    `db:"wallet_id"`
	CountDaily    int16     `db:"count_daily"`
	CountMonthly  int16     `db:"count_monthly"`
	AmountDaily   int64     `db:"amount_daily"`
	AmountMonthly int64     `db:"amount_monthly"`
	UpdatedAt     time.Time `db:"updated_at"`
}
