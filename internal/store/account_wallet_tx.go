package store

import (
	"github.com/jmoiron/sqlx"
)

type AccountWalletStore struct {
	DB *sqlx.DB
}

// func NewAccountWalletStore(db *sqlx.DB) *AccountStore {
// 	return &AccountWalletStore{
// 		DB: db,
// 	}
// }

// func (s *AccountWalletStore) CreateAccountWallet(ctx context.Context, a entity.Account, w entity.Wallet) error {
// 	tx, err := s.DB.Begin()
// 	tx.
// 	return err
// }
