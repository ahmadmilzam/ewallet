package entity

import (
	"context"
	"time"
)

//go:generate mockery --name AccountStoreQuerier
type AccountQuery interface {
	CreateAccount(ctx context.Context, a *Account) (*Account, error)
	CreateAccountTx(ctx context.Context, a *Account, w *Wallet) error
	UpgradeAccount(ctx context.Context, a *Account) (*Account, error)
	FindAccountForUpdateByPhone(ctx context.Context, p string) (*Account, error)
	FindAccountByPhone(ctx context.Context, p string) (*Account, error)
	FindAccountAndWalletsByPhone(ctx context.Context, p string) ([]AccountWallet, error)
}

type Account struct {
	Phone     string    `db:"phone" faker:"customphone,unique"`
	Name      string    `db:"name" faker:"name,unique"`
	Email     string    `db:"email" faker:"email,unique"`
	Role      string    `db:"role" faker:"accountRole"`
	Status    string    `db:"status" faker:"accountStatus"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type AccountWallet struct {
	Account
	WalletSummary
}
