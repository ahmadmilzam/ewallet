package entity

import (
	"context"
	"time"
)

//go:generate mockery --name AccountStoreQuerier
type AccountQuery interface {
	CreateAccount(ctx context.Context, model Account) (Account, error)
	CreateAccountTx(ctx context.Context, a Account, w Wallet) error
	UpgradeAccount(ctx context.Context, id string) (Account, error)
	FindAccountById(ctx context.Context, id string) (Account, error)
	FindAccountByPhone(ctx context.Context, phone string) (Account, error)
}

type Account struct {
	ID        string    `json:"id" db:"id" faker:"uuid_hyphenated,unique"`
	Phone     string    `json:"phone" db:"phone" faker:"customphone,unique"`
	Name      string    `json:"name" db:"name" faker:"name,unique"`
	Email     string    `json:"email" db:"email" faker:"email,unique"`
	Role      string    `json:"role" db:"role" faker:"accountRole"`
	Status    string    `json:"status" db:"status" faker:"accountStatus"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
