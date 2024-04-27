package entity

import (
	"context"
	"time"
)

//go:generate mockery --name JournalStoreQuerier
type TransferQuery interface {
	CreateTransfer(ctx context.Context, transfer *Transfer) (*Transfer, error)
	CreateTransferTx(
		ctx context.Context,
		transfer *Transfer,
		entries []Entry,
		wallets []WalletUpdateBalance,
		counter *UpdateTransferCounter,
	) error
	FindTransferById(ctx context.Context, id string) (*Transfer, error)
}

type Transfer struct {
	ID          string    `db:"id"`
	SrcWalletID string    `db:"src_wallet_id"`
	DstWalletID string    `db:"dst_wallet_id"`
	Amount      int64     `db:"amount"`
	Type        string    `db:"type"`
	Reference   string    `db:"reference"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
