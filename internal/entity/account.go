package entity

import (
	"context"
	"time"
)

type AccountQuery interface {
	CreateAccount(ctx context.Context, account *Account) (*Account, error)
	CreateAccountTx(ctx context.Context, account *Account, wallets []Wallet, counter *TransferCounter) error
	UpdateAccount(ctx context.Context, account *Account) (*Account, error)
	FindAccountForUpdateById(ctx context.Context, id string) (*Account, error)
	FindAccountById(ctx context.Context, id string) (*Account, error)
	FindAccountAndWalletsById(ctx context.Context, id string) ([]AccountWallet, error)
}

type Account struct {
	Phone     string    `db:"phone"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Role      string    `db:"role"`
	Status    string    `db:"status"`
	COAType   string    `db:"coa_type"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type AccountWallet struct {
	Account
	Wallet `db:"wallet"`
}
