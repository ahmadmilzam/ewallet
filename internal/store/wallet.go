package store

import (
	"context"
	"fmt"

	"github.com/ahmadmilzam/ewallet/internal/entity"
)

var (
	CreateWalletSQL = `INSERT INTO wallets VALUES(:id, :account_phone, :balance, :type, :created_at, :updated_at)`
)

func (s *Queries) CreateWallet(ctx context.Context, w *entity.Wallet) (*entity.Wallet, error) {
	_, err := s.db.NamedExecContext(ctx, CreateWalletSQL, w)

	if err != nil {
		return nil, fmt.Errorf("error creating wallet: %w", err)
	}

	return w, nil
}

func (s *Queries) FindWalletById(ctx context.Context, id string) (*entity.Wallet, error) {
	fmt.Println("store/wallet")
	var mw entity.Wallet
	err := s.db.GetContext(ctx, &mw, `SELECT * FROM wallets WHERE id = $1 LIMIT 1`, id)
	if err != nil {
		return nil, fmt.Errorf("error getting wallet: %w", err)
	}

	return &mw, nil
}

func (s *Queries) FindWalletsByPhone(ctx context.Context, p string) ([]entity.Wallet, error) {
	var amw []entity.Wallet

	err := s.db.SelectContext(ctx, &amw, `SELECT * FROM wallets WHERE account_phone=$1`, p)
	if err != nil {
		return nil, fmt.Errorf("error getting wallets: %w", err)
	}

	return amw, nil
}

func (s *Queries) FindWalletsByPhoneAndType(ctx context.Context, p string, t string) (*entity.Wallet, error) {
	var w entity.Wallet

	err := s.db.SelectContext(ctx, &w, `SELECT * FROM wallets WHERE account_phone=$1 AND type=$2`, p, t)
	if err != nil {
		return nil, fmt.Errorf("error getting wallets: %w", err)
	}

	return &w, nil
}
