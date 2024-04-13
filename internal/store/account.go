package store

import (
	"context"
	"fmt"

	"github.com/ahmadmilzam/ewallet/internal/entity"
)

const (
	createAccountSQL = `INSERT INTO accounts
	VALUES(:phone, :name, :email, :role, :status, :created_at, :updated_at)`
	updateAccountSQL = `UPDATE accounts SET status = 'PREMIUM', updated_at = :updated_at WHERE phone = :phone`
)

func (s *Queries) CreateAccount(ctx context.Context, a *entity.Account) (*entity.Account, error) {
	_, err := s.db.NamedExecContext(ctx, createAccountSQL, a)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (s *SQLStore) CreateAccountTx(ctx context.Context, a *entity.Account, w *entity.Wallet) error {

	err := s.execTx(func(q *Queries) error {
		var err error

		_, err = q.CreateAccount(ctx, a)
		if err != nil {
			return err
		}

		_, err = q.CreateWallet(ctx, w)
		if err != nil {
			return err
		}

		return err
	})

	return err
}

func (s *Queries) UpgradeAccount(ctx context.Context, a *entity.Account) (*entity.Account, error) {
	ac := &entity.Account{}
	_, err := s.db.NamedExecContext(ctx, updateAccountSQL, a)

	if err != nil {
		return nil, fmt.Errorf("error upgrading account: %w", err)
	}

	return ac, nil
}

func (s *Queries) DeleteAccount(ctx context.Context, id string) error {
	_, err := s.db.NamedExecContext(ctx, `DELETE * FROM accounts WHERE id = :id`, id)
	if err != nil {
		return fmt.Errorf("error deleting account: %w", err)
	}

	return nil
}

func (s *Queries) FindAccountForUpdateByPhone(ctx context.Context, p string) (*entity.Account, error) {
	a := &entity.Account{}
	err := s.db.GetContext(ctx, a, `SELECT * FROM accounts WHERE id = $1 LIMIT 1 FOR UPDATE`, p)
	if err != nil {
		return nil, fmt.Errorf("error getting account: %w", err)
	}

	return a, nil
}

func (s *Queries) FindAccountByPhone(ctx context.Context, phone string) (*entity.Account, error) {
	ma := &entity.Account{}
	err := s.db.GetContext(ctx, ma, `SELECT * FROM accounts WHERE phone = $1 LIMIT 1`, phone)
	if err != nil {
		return nil, err
	}

	return ma, nil
}

func (s *Queries) FindAccountAndWalletsByPhone(ctx context.Context, p string) ([]entity.AccountWallet, error) {
	var aaw []entity.AccountWallet
	// var aw entity.AccountWallets

	const q = `
		SELECT
			a.phone,
			a.name,
			a.email,
			a.role,
			a.status,
			a.created_at,
			w.type,
			w.balance
		FROM
			accounts a
		JOIN wallets AS w ON a.phone = w.account_phone
		WHERE
		a.phone = $1
	`
	err := s.db.SelectContext(ctx, &aaw, q, p)
	if err != nil {
		return nil, err
	}

	return aaw, nil
}

func (s *Queries) FindAccounts(ctx context.Context) ([]entity.Account, error) {
	var ama []entity.Account

	err := s.db.SelectContext(ctx, &ama, `SELECT * FROM accounts`)
	if err != nil {
		return []entity.Account{}, fmt.Errorf("error getting accounts: %w", err)
	}

	return ama, nil
}
