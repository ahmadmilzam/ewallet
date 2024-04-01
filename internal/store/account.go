package store

import (
	"context"
	"fmt"

	"github.com/ahmadmilzam/ewallet/internal/model"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type AccountStore struct {
	DB *sqlx.DB
}

func NewAccountStore(db *sqlx.DB) *AccountStore {
	return &AccountStore{
		DB: db,
	}
}

func (s *AccountStore) CreateAccount(ctx context.Context, a *model.Account) (*model.Account, error) {
	var ac model.Account
	err := s.DB.GetContext(ctx, &ac, `INSERT INTO accounts VALUES($1, $2, $3) RETURNING *`,
		a.ID,
		a.Phone,
		a.Status,
	)

	if err != nil {
		return nil, fmt.Errorf("error creating account: %w", err)
	}

	return &ac, nil
}

func (s *AccountStore) UpgradeAccount(ctx context.Context, id uuid.UUID) (*model.Account, error) {
	var ac model.Account
	err := s.DB.GetContext(ctx, &ac, `UPDATE accounts SET status = 'PREMIUM' WHERE id = $1 RETURNING *`, id)

	if err != nil {
		return nil, fmt.Errorf("error upgrading account: %w", err)
	}

	return &ac, nil
}

func (s *AccountStore) DeleteAccount(ctx context.Context, id uuid.UUID) error {
	_, err := s.DB.ExecContext(ctx, `DELETE * FROM accounts WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("error deleting account: %w", err)
	}

	return nil
}

func (s *AccountStore) FindAccount(ctx context.Context, id uuid.UUID) (*model.Account, error) {
	var a model.Account
	err := s.DB.GetContext(ctx, &a, `SELECT * FROM accounts WHERE id = $1 LIMIT 1`, id)
	if err != nil {
		return nil, fmt.Errorf("error getting account: %w", err)
	}

	return &a, nil
}

func (s *AccountStore) FindAccounts(ctx context.Context) ([]model.Account, error) {
	var aa []model.Account

	err := s.DB.SelectContext(ctx, &aa, `SELECT * FROM accounts`)
	if err != nil {
		return nil, fmt.Errorf("error getting accounts: %w", err)
	}

	return aa, nil
}
