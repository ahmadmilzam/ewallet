package store

import (
	"context"
	"fmt"

	"github.com/ahmadmilzam/ewallet/internal/entity"
)

var (
	CreateCounterSQL = `INSERT INTO transfer_counters
	VALUES(:wallet_id, :credit_count_daily, :credit_count_monthly, :credit_amount_daily, :credit_amount_monthly, :created_at, :updated_at)`
	UpdateCounterSQL = `
	UPDATE transfer_counters
	SET
		credit_count_daily = credit_count_daily + :count_daily,
		credit_count_monthly = credit_count_monthly + :count_monthly,
		credit_amount_daily = credit_amount_daily + :amount_daily,
		credit_amount_monthly = credit_amount_monthly + :amount_monthly,
		updated_at = :updated_at
	WHERE wallet_id = :wallet_id`
	FindCounterByIdSQL          = `SELECT * FROM transfer_counters WHERE wallet_id = $1 LIMIT 1`
	FindCounterForUpdateByIdSQL = `SELECT * FROM transfer_counters WHERE wallet_id = $1 LIMIT 1 FOR UPDATE`
)

func (s *Queries) CreateCounter(ctx context.Context, counter *entity.TransferCounter) (*entity.TransferCounter, error) {
	_, err := s.db.NamedExecContext(ctx, CreateCounterSQL, counter)

	if err != nil {
		return nil, fmt.Errorf("CreateCounter: %w", err)
	}

	return counter, nil
}

func (s *Queries) UpdateCounter(ctx context.Context, counter *entity.UpdateTransferCounter) error {
	_, err := s.db.NamedExecContext(ctx, UpdateCounterSQL, counter)

	if err != nil {
		return fmt.Errorf("UpdateCounter: %w", err)
	}

	return nil
}

func (s *Queries) FindCounterById(ctx context.Context, id string) (*entity.TransferCounter, error) {
	counter := &entity.TransferCounter{}
	err := s.db.GetContext(ctx, counter, FindCounterByIdSQL, id)
	fmt.Println("Counter data: ", counter)
	if err != nil {
		return nil, fmt.Errorf("FindCounterById: %w", err)
	}

	return counter, nil
}

func (s *Queries) FindCounterForUpdateById(ctx context.Context, id string) (*entity.TransferCounter, error) {
	var tc *entity.TransferCounter
	err := s.db.GetContext(ctx, tc, FindCounterForUpdateByIdSQL, id)
	if err != nil {
		return nil, fmt.Errorf("FindCounterById: %w", err)
	}

	return tc, nil
}
