package store

import (
	"context"
	"fmt"

	"github.com/ahmadmilzam/ewallet/internal/entity"
)

const (
	CreateTransferSQL = `
	INSERT INTO transfers
	VALUES(
		:id,
		:src_wallet_id,
		:dst_wallet_id,
		:amount,
		:type,
		:reference,
		:created_at,
		:updated_at
	)`
	UpdateTransferSQL = `UPDATE transfers
	SET id = :id, src_wallet_id = :src_wallet_id, dst_wallet_id = :dst_wallet_id, amount = :amount, reference = :reference, type = :type, updated_at = :updated_at
	WHERE id = :id`
	FindTransferByIdSQL          = `SELECT * FROM transfers WHERE id = $1 LIMIT 1`
	FindTransferForUpdateByIdSQL = `SELECT * FROM transfers WHERE id = $1 LIMIT 1 FOR UPDATE`
)

func (s *QueryCommands) CreateTransfer(ctx context.Context, transfer *entity.Transfer) (*entity.Transfer, error) {
	_, err := s.db.NamedExecContext(ctx, CreateTransferSQL, transfer)

	if err != nil {
		return nil, fmt.Errorf("CreateTransfer: %w", err)
	}

	return transfer, nil
}

func (s *QueryCommands) FindTransferById(ctx context.Context, id string) (*entity.Transfer, error) {
	counter := &entity.Transfer{}
	err := s.db.GetContext(ctx, counter, FindTransferByIdSQL, id)
	if err != nil {
		return nil, fmt.Errorf("FindTransferById: %w", err)
	}

	return counter, nil
}
