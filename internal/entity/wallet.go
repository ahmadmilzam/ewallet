package entity

import (
	"context"
	"time"
)

//go:generate mockery --name WalletStoreQuerier
type WalletQuery interface {
	CreateWallet(ctx context.Context, wallet *Wallet) (*Wallet, error)
	UpdateWallet(ctx context.Context, wallet *Wallet) error
	UpdateWalletBalance(ctx context.Context, wallet *WalletUpdateBalance) error
	FindWalletById(ctx context.Context, id string) (*Wallet, error)
	FindWalletForUpdateById(ctx context.Context, id string) (*Wallet, error)
	FindWalletsByPhone(ctx context.Context, p string) ([]Wallet, error)
}

type Wallet struct {
	ID           string    `db:"id"`
	AccountPhone string    `db:"account_phone"`
	Balance      int64     `db:"balance"`
	Type         string    `db:"type"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type WalletSummary struct {
	Balance int64  `db:"balance"`
	Type    string `db:"type"`
}

type WalletUpdateBalance struct {
	ID        string    `db:"id"`
	Amount    int64     `db:"amount"`
	UpdatedAt time.Time `db:"updated_at"`
}
