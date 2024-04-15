package store

import (
	"context"
	"fmt"

	"github.com/ahmadmilzam/ewallet/internal/entity"
)

var (
	CreateCounterSQL = `INSERT INTO transfer_counters
	VALUES(:wallet_id, :count_daily, :count_monthly, :amount_daily, :amount_monthly, :created_at, :updated_at)`
	UpdateCounterSQL = `UPDATE transfer_counters
	SET wallet_id = :wallet_id, count_daily = :count_daily, count_monthly = :count_monthly, amount_daily = :amount_daily, amount_monthly = :amount_monthly, updated_at = :updated_at WHERE wallet_id = :wallet_id`
	FindCounterByIdSQL          = `SELECT * FROM transfer_counters WHERE id = $1 LIMIT 1`
	FindCounterForUpdateByIdSQL = `SELECT * FROM transfer_counters WHERE id = $1 LIMIT 1 FOR UPDATE`
)

func (s *Queries) CreateCounter(ctx context.Context, w *entity.TransferCounter) (*entity.TransferCounter, error) {
	_, err := s.db.NamedExecContext(ctx, CreateCounterSQL, w)

	if err != nil {
		return nil, fmt.Errorf("CreateCounter: %w", err)
	}

	return w, nil
}

func (s *Queries) UpdateCounter(ctx context.Context, tc *entity.TransferCounter) (*entity.TransferCounter, error) {
	_, err := s.db.NamedExecContext(ctx, UpdateCounterSQL, tc)

	if err != nil {
		return nil, fmt.Errorf("UpdateCounter: %w", err)
	}

	return tc, nil
}

func (s *Queries) FindCounterById(ctx context.Context, id string) (*entity.TransferCounter, error) {
	var tc *entity.TransferCounter
	err := s.db.GetContext(ctx, tc, FindCounterByIdSQL, id)
	if err != nil {
		return nil, fmt.Errorf("FindCounterById: %w", err)
	}

	return tc, nil
}

func (s *Queries) FindCounterForUpdateById(ctx context.Context, id string) (*entity.TransferCounter, error) {
	var tc *entity.TransferCounter
	err := s.db.GetContext(ctx, tc, FindCounterForUpdateByIdSQL, id)
	if err != nil {
		return nil, fmt.Errorf("FindCounterById: %w", err)
	}

	return tc, nil
}
