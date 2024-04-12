package store

import (
	"context"
	"fmt"

	"github.com/ahmadmilzam/ewallet/internal/entity"
)

func (s *Queries) CreateWallet(ctx context.Context, w entity.Wallet) (*entity.Wallet, error) {
	_, err := s.db.Exec(`INSERT INTO wallets VALUES($1, $2, $3, $4, $5, $6)`,
		w.ID,
		w.AccountId,
		w.Balance,
		w.Type,
		w.CreatedAt,
		w.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error creating wallet: %w", err)
	}

	return &w, nil
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

func (s *Queries) FindWalletsByAccount(ctx context.Context, aid string) ([]entity.Wallet, error) {
	var amw []entity.Wallet

	err := s.db.SelectContext(ctx, &amw, `SELECT * FROM wallets WHERE account_id=$1`, aid)
	if err != nil {
		return []entity.Wallet{}, fmt.Errorf("error getting wallets: %w", err)
	}

	return amw, nil
}
