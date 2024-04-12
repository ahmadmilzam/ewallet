package entity

import (
	"context"
	"time"
)

//go:generate mockery --name WalletStoreQuerier
type WalletQuery interface {
	CreateWallet(ctx context.Context, model Wallet) (*Wallet, error)
	FindWalletById(ctx context.Context, id string) (*Wallet, error)
	FindWalletsByAccount(ctx context.Context, aid string) ([]Wallet, error)
}

type Wallet struct {
	ID        string    `json:"id" db:"id"`
	AccountId string    `json:"account_id" db:"account_id"`
	Balance   float64   `json:"balance" db:"balance"`
	Type      string    `json:"type" db:"type"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
