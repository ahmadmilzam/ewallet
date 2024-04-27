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

func (s *Queries) CreateTransfer(ctx context.Context, transfer *entity.Transfer) (*entity.Transfer, error) {
	_, err := s.db.NamedExecContext(ctx, CreateTransferSQL, transfer)

	if err != nil {
		return nil, fmt.Errorf("CreateTransfer: %w", err)
	}

	return transfer, nil
}

func (s *SQLStore) CreateTransferTx(ctx context.Context, transfer *entity.Transfer, entries []entity.Entry, wallets []entity.WalletUpdateBalance, counter *entity.UpdateTransferCounter) error {

	err := s.execTx(func(q *Queries) error {
		var err error
		// if needSrcLock {
		// 	_, err = q.FindWalletForUpdateById(ctx, wallets["src"].ID)
		// 	if err != nil {
		// 		err = fmt.Errorf("CreateTransferTx: %w", err)
		// 		return err
		// 	}
		// }

		// if needDstLock {
		// 	_, err = q.FindWalletForUpdateById(ctx, wallets["dst"].ID)
		// 	if err != nil {
		// 		err = fmt.Errorf("CreateTransferTx: %w", err)
		// 		return err
		// 	}
		// }

		_, err = q.CreateTransfer(ctx, transfer)
		if err != nil {
			err = fmt.Errorf("CreateTransferTx: %w", err)
			return err
		}

		for _, entry := range entries {
			_, err = q.CreateEntry(ctx, &entry)
			if err != nil {
				err = fmt.Errorf("CreateTransferTx: %w", err)
				return err
			}
		}

		for _, wallet := range wallets {
			err = q.UpdateWalletBalance(ctx, &wallet)
			if err != nil {
				err = fmt.Errorf("CreateTransferTx: %w", err)
				return err
			}
		}

		err = q.UpdateCounter(ctx, counter)
		if err != nil {
			err = fmt.Errorf("CreateTransferTx: %w", err)
			return err
		}

		return err
	})

	return err
}

func (s *Queries) FindTransferById(ctx context.Context, id string) (*entity.Transfer, error) {
	counter := &entity.Transfer{}
	err := s.db.GetContext(ctx, counter, FindTransferByIdSQL, id)
	if err != nil {
		return nil, fmt.Errorf("FindTransferById: %w", err)
	}

	return counter, nil
}

// func (s *Queries) UpdateTransfer(ctx context.Context, counter *entity.Transfer) (*entity.Transfer, error) {
// 	_, err := s.db.NamedExecContext(ctx, UpdateTransferSQL, counter)

// 	if err != nil {
// 		return nil, fmt.Errorf("UpdateCounter: %w", err)
// 	}

// 	return counter, nil
// }

// func (s *Queries) FindTransferForUpdateById(ctx context.Context, id string) (*entity.Transfer, error) {
// 	var tc *entity.Transfer
// 	err := s.db.GetContext(ctx, tc, FindTransferForUpdateByIdSQL, id)
// 	if err != nil {
// 		return nil, fmt.Errorf("FindCounterById: %w", err)
// 	}

// 	return tc, nil
// }
