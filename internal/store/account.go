package store

import (
	"context"
	"fmt"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/google/uuid"
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

func (as *AccountStore) CreateAccount(ctx context.Context, a entity.Account) (entity.Account, error) {
	var ma entity.Account
	err := as.DB.GetContext(ctx, &ma, `INSERT INTO accounts VALUES($1, $2, $3) RETURNING *`,
		a.ID,
		a.Phone,
		a.Status,
	)

	if err != nil {
		return entity.Account{}, fmt.Errorf("error creating account: %w", err)
	}

	return ma, nil
}

func (as *AccountStore) UpgradeAccount(ctx context.Context, id uuid.UUID) (entity.Account, error) {
	var ac entity.Account
	err := as.DB.GetContext(ctx, &ac, `UPDATE accounts SET status = 'PREMIUM' WHERE id = $1 RETURNING *`, id)

	if err != nil {
		return entity.Account{}, fmt.Errorf("error upgrading account: %w", err)
	}

	return ac, nil
}

func (as *AccountStore) DeleteAccount(ctx context.Context, id uuid.UUID) error {
	_, err := as.DB.ExecContext(ctx, `DELETE * FROM accounts WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("error deleting account: %w", err)
	}

	return nil
}

func (as *AccountStore) FindAccountById(ctx context.Context, id uuid.UUID) (entity.Account, error) {
	var ma entity.Account
	err := as.DB.GetContext(ctx, &ma, `SELECT * FROM accounts WHERE id = $1 LIMIT 1`, id)
	if err != nil {
		return entity.Account{}, fmt.Errorf("error getting account: %w", err)
	}

	return ma, nil
}

func (as *AccountStore) FindAccountByPhone(ctx context.Context, phone string) (entity.Account, error) {
	var ma entity.Account
	err := as.DB.GetContext(ctx, &ma, `SELECT * FROM accounts WHERE phone = $1 LIMIT 1`, phone)
	if err != nil {
		return entity.Account{}, fmt.Errorf("error getting account: %w", err)
	}

	return ma, nil
}

func (as *AccountStore) FindAccounts(ctx context.Context) ([]entity.Account, error) {
	var ama []entity.Account

	err := as.DB.SelectContext(ctx, &ama, `SELECT * FROM accounts`)
	if err != nil {
		return []entity.Account{}, fmt.Errorf("error getting accounts: %w", err)
	}

	return ama, nil
}
