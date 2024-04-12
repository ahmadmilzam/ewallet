package store

import (
	"context"
	"fmt"

	"github.com/ahmadmilzam/ewallet/internal/entity"
)

func (s *Queries) CreateAccount(ctx context.Context, a entity.Account) (entity.Account, error) {
	_, err := s.db.ExecContext(ctx, `INSERT INTO accounts VALUES($1, $2, $3, $4, $5, $6, $7, $8)`,
		a.ID,
		a.Phone,
		a.Name,
		a.Email,
		a.Role,
		a.Status,
		a.CreatedAt,
		a.UpdatedAt,
	)

	if err != nil {
		return entity.Account{}, err
	}

	return a, nil
}

func (s *SQLStore) CreateAccountTx(ctx context.Context, a entity.Account, w entity.Wallet) error {

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

func (s *Queries) UpgradeAccount(ctx context.Context, id string) (entity.Account, error) {
	var ac entity.Account
	err := s.db.GetContext(ctx, &ac, `UPDATE accounts SET status = 'PREMIUM' WHERE id = $1 RETURNING *`, id)

	if err != nil {
		return entity.Account{}, fmt.Errorf("error upgrading account: %w", err)
	}

	return ac, nil
}

func (s *Queries) DeleteAccount(ctx context.Context, id string) error {
	_, err := s.db.ExecContext(ctx, `DELETE * FROM accounts WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("error deleting account: %w", err)
	}

	return nil
}

func (s *Queries) FindAccountById(ctx context.Context, id string) (entity.Account, error) {
	var ma entity.Account
	err := s.db.GetContext(ctx, &ma, `SELECT * FROM accounts WHERE id = $1 LIMIT 1`, id)
	if err != nil {
		return entity.Account{}, fmt.Errorf("error getting account: %w", err)
	}

	return ma, nil
}

func (s *Queries) FindAccountByPhone(ctx context.Context, phone string) (entity.Account, error) {
	var ma entity.Account
	err := s.db.GetContext(ctx, &ma, `SELECT * FROM accounts WHERE phone = $1 LIMIT 1`, phone)
	if err != nil {
		return entity.Account{}, err
	}

	return ma, nil
}

func (s *Queries) FindAccounts(ctx context.Context) ([]entity.Account, error) {
	var ama []entity.Account

	err := s.db.SelectContext(ctx, &ama, `SELECT * FROM accounts`)
	if err != nil {
		return []entity.Account{}, fmt.Errorf("error getting accounts: %w", err)
	}

	return ama, nil
}
