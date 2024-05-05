package store

import (
	"context"
	"fmt"

	"github.com/ahmadmilzam/ewallet/internal/entity"
)

func (s *SQLStore) execTx(fn func(*QueryCommands) error) error {
	tx, err := s.DB.Beginx()
	if err != nil {
		return err
	}

	q := NewQueryCommands(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

func (s *SQLStore) CreateAccountTx(ctx context.Context, account *entity.Account, wallets []entity.Wallet, counter *entity.TransferCounter) error {

	err := s.execTx(func(q *QueryCommands) error {
		var err error

		_, err = q.CreateAccount(ctx, account)
		if err != nil {
			return fmt.Errorf("CreateAccountTx: %w", err)
		}

		_, err = q.CreateWallet(ctx, &wallets[0])
		if err != nil {
			err = fmt.Errorf("CreateAccountTx: %w", err)
			return err
		}

		_, err = q.CreateWallet(ctx, &wallets[1])
		if err != nil {
			return fmt.Errorf("CreateAccountTx: %w", err)
		}

		_, err = q.CreateCounter(ctx, counter)
		if err != nil {
			return fmt.Errorf("CreateAccountTx: %w", err)
		}

		return err
	})

	return err
}

func (s *SQLStore) CreateTransferTx(ctx context.Context, transfer *entity.Transfer, entries []entity.Entry, wallets []entity.WalletUpdateBalance, counter *entity.TransferCounter, lockCounter bool) error {
	err := s.execTx(func(q *QueryCommands) error {
		var err error

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

		if lockCounter {
			_, err = q.FindCounterForUpdateById(ctx, counter.WalletID)
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
