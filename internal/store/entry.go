package store

import (
	"context"
	"fmt"

	"github.com/ahmadmilzam/ewallet/internal/entity"
)

const (
	CreateEntrySQL = `INSERT INTO entries
	VALUES(:id, :wallet_id, :credit_amount, :debit_amount, :balance_before, :balance_after, :correlation_id, :transfer_id, :created_at, :updated_at)`
	UpdateEntrySQL = `UPDATE entries
	SET id = :id, wallet_id = :wallet_id, credit_amount = :credit_amount, debit_amount = :debit_amount, balance_before = :balance_before, balance_after = :balance_after, correlation_id = :correlation_id, transfer_id = :transfer_id, created_at = :created_at, updated_at = :updated_at
	WHERE wallet_id = :wallet_id`
	FindEntryByIdSQL = `SELECT * FROM entries WHERE id = $1 LIMIT 1`
)

func (s *Queries) CreateEntry(ctx context.Context, entry *entity.Entry) (*entity.Entry, error) {
	_, err := s.db.NamedExecContext(ctx, CreateEntrySQL, entry)

	if err != nil {
		return nil, fmt.Errorf("CreateEntry: %w", err)
	}

	return entry, nil
}

func (s *Queries) UpdateEntry(ctx context.Context, entry *entity.Entry) (*entity.Entry, error) {
	_, err := s.db.NamedExecContext(ctx, UpdateEntrySQL, entry)

	if err != nil {
		return nil, fmt.Errorf("UpdateEntry: %w", err)
	}

	return entry, nil
}

func (s *Queries) FindEntryById(ctx context.Context, id string) (*entity.Entry, error) {
	entry := &entity.Entry{}
	err := s.db.GetContext(ctx, entry, FindEntryByIdSQL, id)
	if err != nil {
		return nil, fmt.Errorf("FindEntryById: %w", err)
	}

	return entry, nil
}

// func (s *Queries) FindEntryForUpdateById(ctx context.Context, id string) (*entity.Entry, error) {
// 	var tc *entity.Entry
// 	err := s.db.GetContext(ctx, tc, FindEntryForUpdateByIdSQL, id)
// 	if err != nil {
// 		return nil, fmt.Errorf("FindEntryById: %w", err)
// 	}

// 	return tc, nil
// }
