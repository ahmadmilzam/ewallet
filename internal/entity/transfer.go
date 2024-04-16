package entity

import (
	"context"
	"database/sql"
	"time"
)

//go:generate mockery --name JournalStoreQuerier
type TransferQuery interface {
	CreateTransfer(ctx context.Context, transfer Transfer) (*Transfer, error)
	UpdateTransfer(ctx context.Context, transfer Transfer) (*Transfer, error)
	DeleteTransfer(ctx context.Context, id string) error
	FindTransfer(ctx context.Context, id string) (*Transfer, error)
	FindTransfers(ctx context.Context) ([]Transfer, error)
}

type Transfer struct {
	ID          string         `json:"id" db:"id"`
	SrcWalletID string         `json:"src_wallet_id" db:"src_wallet_id"`
	DstWalletID string         `json:"dst_wallet_id" db:"dst_wallet_id"`
	Amount      float64        `json:"amount" db:"amount"`
	Reference   string         `json:"reference" db:"reference"`
	Type        sql.NullString `json:"type" db:"type"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
}
