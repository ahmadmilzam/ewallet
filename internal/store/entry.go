package store

import (
	"context"
	"fmt"

	"github.com/ahmadmilzam/ewallet/internal/entity"
)

const (
	CreateEntrySQL = `INSERT INTO entries
	VALUES(:id, :wallet_id, :credit_amount, :debit_amount, :balance_before, :balance_after, :correlation_id, :transfer_id, :created_at, :updated_at)`
	FindEntryByIdSQL = `SELECT * FROM entries WHERE id = $1 LIMIT 1`
)

func (s *QueryCommands) CreateEntry(ctx context.Context, entry *entity.Entry) (*entity.Entry, error) {
	_, err := s.db.NamedExecContext(ctx, CreateEntrySQL, entry)

	if err != nil {
		return nil, fmt.Errorf("CreateEntry: %w", err)
	}

	return entry, nil
}

func (s *QueryCommands) FindEntryById(ctx context.Context, id string) (*entity.Entry, error) {
	entry := &entity.Entry{}
	err := s.db.GetContext(ctx, entry, FindEntryByIdSQL, id)
	if err != nil {
		return nil, fmt.Errorf("FindEntryById: %w", err)
	}

	return entry, nil
}
