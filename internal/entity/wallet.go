package entity

import (
	"context"
	"time"
)

//go:generate mockery --name WalletStoreQuerier
type WalletQuery interface {
	CreateWallet(ctx context.Context, model *Wallet) (*Wallet, error)
	FindWalletById(ctx context.Context, id string) (*Wallet, error)
	FindWalletsByPhone(ctx context.Context, p string) ([]Wallet, error)
}

type Wallet struct {
	ID           string    `db:"id"`
	AccountPhone string    `db:"account_phone"`
	Balance      float64   `db:"balance"`
	Type         string    `db:"type"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type WalletSummary struct {
	Balance float64 `db:"balance"`
	Type    string  `db:"type"`
}
