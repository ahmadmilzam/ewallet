package entity

import (
	"context"
	"time"
)

//go:generate mockery --name AccountSQLStore
type AccountQueryStore interface {
	CreateAccount(ctx context.Context, model Account) (Account, error)
	UpgradeAccount(ctx context.Context, id string) (Account, error)
	DeleteAccount(ctx context.Context, id string) error
	FindAccountById(ctx context.Context, id string) (Account, error)
	FindAccountByPhone(ctx context.Context, phone string) (Account, error)
	FindAccounts(ctx context.Context) ([]Account, error)
}

type Account struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Phone     string    `json:"phone" db:"phone"`
	Role      string    `json:"role" db:"role"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type WalletQueryStore interface {
	CreateWallet(ctx context.Context, model Wallet) (Wallet, error)
	DeleteWallet(ctx context.Context, id string) error
	FindWallet(ctx context.Context, id string) (Wallet, error)
	FindAccountWallets(ctx context.Context, wid string) ([]Wallet, error)
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
	DeleteTransfer(ctx context.Context, id string) error
	FindTransfer(ctx context.Context, id string) (Transfer, error)
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
	DeleteJournal(ctx context.Context, id string) error
	FindJournal(ctx context.Context, id string) (Journal, error)
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
