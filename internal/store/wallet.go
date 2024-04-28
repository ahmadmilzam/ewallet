package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/lib/pq"
)

const (
	CreateWalletSQL = `INSERT INTO wallets VALUES(:id, :account_phone, :balance, :type, :created_at, :updated_at)`
	UpdateWalletSQL = `
	UPDATE wallets
	SET
		id = :id,
		account_phone = :account_phone,
		balance = :balance,
		type = :type,
		created_at = :created_at,
		updated_at = :updated_at
	WHERE id = :id`
	UpdateWalletBalaceSQL = `
	UPDATE wallets
	SET
		balance = balance + :amount,
		updated_at = :updated_at
	WHERE id = :id`
	FindWalletByIdSQL          = `SELECT * FROM wallets WHERE id = $1 LIMIT 1`
	FindWalletForUpdateByIdSQL = `SELECT * FROM wallets WHERE id = $1 LIMIT 1 FOR UPDATE`
	FindWalletByPhoneSQL       = `SELECT * FROM wallets WHERE account_phone=$1`
)

func (s *Queries) CreateWallet(ctx context.Context, w *entity.Wallet) (*entity.Wallet, error) {
	_, err := s.db.NamedExecContext(ctx, CreateWalletSQL, w)

	if err != nil {
		return nil, fmt.Errorf("CreateWallet: %w", err)
	}

	return w, nil
}

func (s *Queries) UpdateWallet(ctx context.Context, wallet *entity.Wallet) error {
	results, err := s.db.NamedExecContext(ctx, UpdateWalletSQL, wallet)
	if err != nil {
		return fmt.Errorf("UpdateWallet: %w", err)
	}

	affected, _ := results.RowsAffected()
	if affected <= 0 {
		return errors.New("UpdateWallet: fail, no rows updated")
	}

	return nil
}

func (s *Queries) UpdateWalletBalance(ctx context.Context, wallet *entity.WalletUpdateBalance) error {

	results, err := s.db.NamedExecContext(ctx, UpdateWalletBalaceSQL, wallet)
	if err != nil {
		return fmt.Errorf("UpdateWalletBalance: %w", err)
	}

	affected, err := results.RowsAffected()
	if err != nil {
		return errors.New("UpdateWalletBalance: RowsAffected: fail, unsupported operation")
	}
	if affected <= 0 {
		return errors.New("UpdateWalletBalance: fail, no rows updated")
	}

	return nil
}

func (s *Queries) FindWalletById(ctx context.Context, id string) (*entity.Wallet, error) {
	var wallet = &entity.Wallet{}
	err := s.db.GetContext(ctx, wallet, FindWalletByIdSQL, id)

	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func (s *Queries) FindWalletForUpdateById(ctx context.Context, id string) (*entity.Wallet, error) {
	var mw entity.Wallet
	err := s.db.GetContext(ctx, &mw, FindWalletForUpdateByIdSQL, id)
	if err != nil {
		return nil, fmt.Errorf("FindWalletForUpdateById: %w", err)
	}

	return &mw, nil
}

func (s *Queries) FindWalletsByPhone(ctx context.Context, phone string) ([]entity.Wallet, error) {
	var wallets []entity.Wallet

	err := s.db.SelectContext(ctx, &wallets, FindWalletByPhoneSQL, phone)

	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			fmt.Println("pq error:", err)
			fmt.Println("pq error:", err.Code.Name())
			/*
				pq error: pq: duplicate key value violates unique constraint "accounts_pkey"
				pq error: unique_violation
			*/
			switch err.Code.Name() {
			case "unique_violation":
				return nil, ErrUniqueViolation

			}
		}

		fmt.Println("non pq error:", err)
		return nil, err
	}

	return wallets, nil
}
