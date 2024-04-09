package store

import (
	"context"
	"fmt"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/jmoiron/sqlx"
)

type WalletStore struct {
	DB *sqlx.DB
}

func NewWalletStore(db *sqlx.DB) *WalletStore {
	return &WalletStore{
		DB: db,
	}
}

func (ws *WalletStore) CreateWallet(ctx context.Context, w entity.Wallet) (entity.Wallet, error) {
	var mw entity.Wallet
	err := ws.DB.GetContext(ctx, &mw, `INSERT INTO accounts VALUES($1, $2, $3) RETURNING *`,
		w.ID,
		w.AccountId,
		w.Balance,
		w.Type,
	)

	if err != nil {
		return entity.Wallet{}, fmt.Errorf("error creating wallet: %w", err)
	}

	return mw, nil
}

func (ws *WalletStore) DeleteWallet(ctx context.Context, id string) error {
	_, err := ws.DB.ExecContext(ctx, `DELETE * FROM wallets WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("error deleting wallet: %w", err)
	}

	return nil
}

func (ws *WalletStore) FindWallet(ctx context.Context, id string) (entity.Wallet, error) {
	var mw entity.Wallet
	err := ws.DB.GetContext(ctx, &mw, `SELECT * FROM wallets WHERE id = $1 LIMIT 1`, id)
	if err != nil {
		return entity.Wallet{}, fmt.Errorf("error getting wallet: %w", err)
	}

	return mw, nil
}

func (ws *WalletStore) FindAccountWallets(ctx context.Context, aid string) ([]entity.Wallet, error) {
	var amw []entity.Wallet

	err := ws.DB.SelectContext(ctx, &amw, `SELECT * FROM wallets WHERE account_id=$1`, aid)
	if err != nil {
		return []entity.Wallet{}, fmt.Errorf("error getting wallets: %w", err)
	}

	return amw, nil
}
