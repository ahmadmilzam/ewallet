package model

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
)

//go:generate mockery --name AccountSQLStore
type AccountStore interface {
	CreateAccount(ctx context.Context, model *Account) (*Account, error)
	UpgradeAccount(ctx context.Context, id uuid.UUID) (*Account, error)
	DeleteAccount(ctx context.Context, id uuid.UUID) error
	FindAccount(ctx context.Context, id uuid.UUID) (*Account, error)
	FindAccounts(ctx context.Context) ([]Account, error)
}

type WalletStore interface {
	CreateWallet(ctx context.Context, model *Wallet) (*Wallet, error)
	UpdateWallet(ctx context.Context, model *Wallet) (*Wallet, error)
	DeleteWallet(ctx context.Context, id uuid.UUID) error
	FindWallet(ctx context.Context, id uuid.UUID) (*Wallet, error)
	FindAccountWallets(ctx context.Context, wid uuid.UUID) ([]*Wallet, error)
}

type TransferStore interface {
	CreateTransfer(ctx context.Context, model *Transfer) (*Transfer, error)
	UpdateTransfer(ctx context.Context, model *Transfer) (*Transfer, error)
	DeleteTransfer(ctx context.Context, id uuid.UUID) error
	FindTransfer(ctx context.Context, id uuid.UUID) (*Transfer, error)
	FindTransfers(ctx context.Context) ([]*Wallet, error)
}

type JournalStore interface {
	CreateJournal(ctx context.Context, model *Journal) (*Journal, error)
	UpdateJournal(ctx context.Context, model *Journal) (*Journal, error)
	DeleteJournal(ctx context.Context, id uuid.UUID) error
	FindJournal(ctx context.Context, id uuid.UUID) (*Journal, error)
	FindJournals(ctx context.Context) ([]*Wallet, error)
}

type Wallet struct {
	ID        uuid.UUID `db:"id"`
	AccountId uuid.UUID `db:"account_id"`
	Balance   int64     `db:"balance"`
	Type      string    `db:"type"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Account struct {
	ID        uuid.UUID `db:"id"`
	Phone     string    `db:"phone"`
	Role      string    `db:"role"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Transfer struct {
	ID            uuid.UUID `db:"id"`
	WalletID      uuid.UUID `db:"wallet_id"`
	CreditAmount  int64     `db:"credit_amount"`
	DebitAmount   int64     `db:"debit_amount"`
	CorrelationID string    `db:"correlation_id"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

type Journal struct {
	ID          uuid.UUID `db:"id"`
	SrcWalletID uuid.UUID `db:"src_wallet_id"`
	DstWalletID uuid.UUID `db:"src_wallet_id"`
	Amount      int64     `db:"amount"`
	Reference   string    `db:"reference"`
	CreatedAt   time.Time `db:"created_at"`
}

type Store interface {
	AccountStore
}
