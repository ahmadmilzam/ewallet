package entity

import (
	"context"
	"time"
)

//go:generate mockery --name JournalStoreQuerier
type JournalQuery interface {
	CreateJournal(ctx context.Context, model Journal) (Journal, error)
	UpdateJournal(ctx context.Context, model Journal) (Journal, error)
	DeleteJournal(ctx context.Context, id string) error
	FindJournal(ctx context.Context, id string) (Journal, error)
	FindJournals(ctx context.Context) ([]Wallet, error)
}

type Journal struct {
	ID          string    `json:"id" db:"id"`
	SrcWalletID string    `json:"src_wallet_id" db:"src_wallet_id"`
	DstWalletID string    `json:"dst_wallet_id" db:"dst_wallet_id"`
	Amount      float64   `json:"amount" db:"amount"`
	Reference   string    `json:"reference" db:"reference"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
