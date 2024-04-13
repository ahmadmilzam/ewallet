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
	FindWalletsByPhoneAndType(ctx context.Context, p string, t string) (*Wallet, error)
}

type Wallet struct {
	ID           string    `json:"id" db:"id"`
	AccountPhone string    `json:"account_phone" db:"account_phone"`
	Balance      float64   `json:"balance" db:"balance"`
	Type         string    `json:"type" db:"type"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type WalletSummary struct {
	Balance float64 `json:"balance" db:"balance"`
	Type    string  `json:"type" db:"type"`
}
