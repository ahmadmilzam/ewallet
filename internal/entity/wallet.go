package entity

import (
	"context"
	"time"
)

//go:generate mockery --name WalletStoreQuerier
type WalletQuery interface {
	CreateWallet(ctx context.Context, wallet *Wallet) (*Wallet, error)
	UpdateWallet(ctx context.Context, wallet *Wallet) error
	FindWalletById(ctx context.Context, id string) (*Wallet, error)
	FindWalletForUpdateById(ctx context.Context, id string) (*Wallet, error)
	FindWalletsByPhone(ctx context.Context, p string) ([]Wallet, error)
}

type Wallet struct {
	ID           string    `db:"id,prefix=wallet_"`
	AccountPhone string    `db:"account_phone"`
	Balance      float64   `db:"balance"`
	Type         string    `db:"type"`
	CreatedAt    time.Time `db:"created_at,prefix=wallet_"`
	UpdatedAt    time.Time `db:"updated_at,prefix=wallet_"`
}

type WalletSummary struct {
	Balance float64 `db:"balance"`
	Type    string  `db:"type"`
}
