package store

import (
	"context"
	"fmt"

	"github.com/ahmadmilzam/ewallet/internal/entity"
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

func (s *AccountStore) CreateAccount(ctx context.Context, a entity.Account) (entity.Account, error) {
	var ma entity.Account
	err := s.DB.GetContext(ctx, &ma, `INSERT INTO accounts VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING *`,
		a.ID,
		a.Phone,
		a.Name,
		a.Role,
		a.Status,
		a.CreatedAt,
		a.UpdatedAt,
	)

	if err != nil {
		return entity.Account{}, err
	}

	return ma, nil
}

func (s *AccountStore) UpgradeAccount(ctx context.Context, id string) (entity.Account, error) {
	var ac entity.Account
	err := s.DB.GetContext(ctx, &ac, `UPDATE accounts SET status = 'PREMIUM' WHERE id = $1 RETURNING *`, id)

	if err != nil {
		return entity.Account{}, fmt.Errorf("error upgrading account: %w", err)
	}

	return ac, nil
}

func (s *AccountStore) DeleteAccount(ctx context.Context, id string) error {
	_, err := s.DB.ExecContext(ctx, `DELETE * FROM accounts WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("error deleting account: %w", err)
	}

	return nil
}

func (s *AccountStore) FindAccountById(ctx context.Context, id string) (entity.Account, error) {
	var ma entity.Account
	err := s.DB.GetContext(ctx, &ma, `SELECT * FROM accounts WHERE id = $1 LIMIT 1`, id)
	if err != nil {
		return entity.Account{}, fmt.Errorf("error getting account: %w", err)
	}

	return ma, nil
}

func (s *AccountStore) FindAccountByPhone(ctx context.Context, phone string) (entity.Account, error) {
	var ma entity.Account
	err := s.DB.GetContext(ctx, &ma, `SELECT * FROM accounts WHERE phone = $1 LIMIT 1`, phone)
	if err != nil {
		return entity.Account{}, err
	}

	return ma, nil
}

func (s *AccountStore) FindAccounts(ctx context.Context) ([]entity.Account, error) {
	var ama []entity.Account

	err := s.DB.SelectContext(ctx, &ama, `SELECT * FROM accounts`)
	if err != nil {
		return []entity.Account{}, fmt.Errorf("error getting accounts: %w", err)
	}

	return ama, nil
}
