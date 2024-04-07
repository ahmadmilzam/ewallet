package entity

import (
	"context"
	"time"

	"github.com/google/uuid"
)

//go:generate mockery --name AccountSQLStore
type AccountQueryStore interface {
	CreateAccount(ctx context.Context, model Account) (Account, error)
	UpgradeAccount(ctx context.Context, id uuid.UUID) (Account, error)
	DeleteAccount(ctx context.Context, id uuid.UUID) error
	FindAccountById(ctx context.Context, id uuid.UUID) (Account, error)
	FindAccountByPhone(ctx context.Context, phone string) (Account, error)
	FindAccounts(ctx context.Context) ([]Account, error)
}

type Account struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	Phone     string    `db:"phone"`
	Role      string    `db:"role"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type WalletQueryStore interface {
	CreateWallet(ctx context.Context, model Wallet) (Wallet, error)
	DeleteWallet(ctx context.Context, id uuid.UUID) error
	FindWallet(ctx context.Context, id uuid.UUID) (Wallet, error)
	FindAccountWallets(ctx context.Context, wid uuid.UUID) ([]Wallet, error)
}

type Wallet struct {
	ID        string    `db:"id"`
	AccountId string    `db:"account_id"`
	Balance   int64     `db:"balance"`
	Type      string    `db:"type"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type TransferQueryStore interface {
	CreateTransfer(ctx context.Context, model Transfer) (Transfer, error)
	UpdateTransfer(ctx context.Context, model Transfer) (Transfer, error)
	DeleteTransfer(ctx context.Context, id uuid.UUID) error
	FindTransfer(ctx context.Context, id uuid.UUID) (Transfer, error)
	FindTransfers(ctx context.Context) ([]Transfer, error)
}

type Transfer struct {
	ID            string    `db:"id"`
	WalletID      string    `db:"wallet_id"`
	CreditAmount  int64     `db:"credit_amount"`
	DebitAmount   int64     `db:"debit_amount"`
	CorrelationID string    `db:"correlation_id"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

type JournalQueryStore interface {
	CreateJournal(ctx context.Context, model Journal) (Journal, error)
	UpdateJournal(ctx context.Context, model Journal) (Journal, error)
	DeleteJournal(ctx context.Context, id uuid.UUID) error
	FindJournal(ctx context.Context, id uuid.UUID) (Journal, error)
	FindJournals(ctx context.Context) ([]Wallet, error)
}

type Journal struct {
	ID          string    `db:"id"`
	SrcWalletID string    `db:"src_wallet_id"`
	DstWalletID string    `db:"src_wallet_id"`
	Amount      int64     `db:"amount"`
	Reference   string    `db:"reference"`
	CreatedAt   time.Time `db:"created_at"`
}

type QueryStore interface {
	AccountQueryStore
}
