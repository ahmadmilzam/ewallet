package store

import (
	"context"
	"fmt"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/lib/pq"
)

const (
	createAccountSQL = `
	INSERT INTO accounts
	VALUES(:phone, :name, :email, :role, :status, :created_at, :updated_at)`
	updateAccountSQL = `
	UPDATE accounts
	SET
		phone = :phone,
		name = :name,
		email = :email,
		role = :role,
		status = :status,
		updated_at = :updated_at
	WHERE phone = :phone`
	deleteAccountByIdSQL         = `DELETE * FROM accounts WHERE phone = :phone`
	findAccountForUpdateByIdSQL  = `SELECT * FROM accounts WHERE phone = $1 LIMIT 1 FOR UPDATE`
	findAccountAndWalletsByIdSQL = `
	SELECT
		a.phone,
		a.name,
		a.email,
		a.role,
		a.status,
		a.created_at,
		a.updated_at,
		w.type,
		w.balance
	FROM
		accounts a
	JOIN wallets AS w ON a.phone = w.account_phone
	WHERE a.phone = $1`
)

func (s *Queries) CreateAccount(ctx context.Context, a *entity.Account) (*entity.Account, error) {
	_, err := s.db.NamedExecContext(ctx, createAccountSQL, a)
	if err, ok := err.(*pq.Error); ok {
		// Here err is of type *pq.Error, you may inspect all its fields, e.g.:
		fmt.Println("pq error:", err)
		fmt.Println("pq error:", err.Code.Name())
		/*
			pq error: pq: duplicate key value violates unique constraint "accounts_pkey"
			pq error: unique_violation
		*/
	}
	if err != nil {
		err = fmt.Errorf("CreateAccount: %w", err)
		return nil, err
	}

	return a, nil
}

func (s *SQLStore) CreateAccountTx(ctx context.Context, a *entity.Account, ww []entity.Wallet, tc *entity.TransferCounter) error {

	err := s.execTx(func(q *Queries) error {
		var err error

		_, err = q.CreateAccount(ctx, a)
		if err != nil {
			err = fmt.Errorf("CreateAccountTx: %w", err)
			return err
		}

		_, err = q.CreateWallet(ctx, &ww[0])
		if err != nil {
			err = fmt.Errorf("CreateAccountTx: %w", err)
			return err
		}

		_, err = q.CreateWallet(ctx, &ww[1])
		if err != nil {
			err = fmt.Errorf("CreateAccountTx: %w", err)
			return err
		}

		_, err = q.CreateCounter(ctx, tc)
		if err != nil {
			err = fmt.Errorf("CreateAccountTx: %w", err)
			return err
		}

		return err
	})

	return err
}

func (s *Queries) UpdateAccount(ctx context.Context, a *entity.Account) (*entity.Account, error) {
	_, err := s.db.NamedExecContext(ctx, updateAccountSQL, a)

	if err != nil {
		return nil, fmt.Errorf("UpdateAccount: %w", err)
	}

	return a, nil
}

func (s *Queries) DeleteAccount(ctx context.Context, id string) error {
	_, err := s.db.NamedExecContext(ctx, deleteAccountByIdSQL, id)
	if err != nil {
		return fmt.Errorf("DeleteAccount: %w", err)
	}

	return nil
}

func (s *Queries) FindAccountForUpdateById(ctx context.Context, p string) (*entity.Account, error) {
	a := &entity.Account{}
	err := s.db.GetContext(ctx, a, findAccountForUpdateByIdSQL, p)
	if err != nil {
		return nil, fmt.Errorf("FindAccountForUpdateByPhone: %w", err)
	}

	return a, nil
}

func (s *Queries) FindAccountById(ctx context.Context, phone string) (*entity.Account, error) {
	ma := &entity.Account{}
	err := s.db.GetContext(ctx, ma, `SELECT * FROM accounts WHERE phone = $1 LIMIT 1`, phone)
	if err != nil {
		err = fmt.Errorf("FindAccountByPhone: %w", err)
		return nil, err
	}

	return ma, nil
}

func (s *Queries) FindAccountAndWalletsById(ctx context.Context, p string) ([]entity.AccountWallet, error) {
	var aaw []entity.AccountWallet

	err := s.db.SelectContext(ctx, &aaw, findAccountAndWalletsByIdSQL, p)
	if err != nil {
		err = fmt.Errorf("FindAccountAndWalletsById: %w", err)
		return nil, err
	}

	return aaw, nil
}
