package store_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/pkg/uuid"
	"github.com/stretchr/testify/assert"
)

func TestStoreTransfer_CreateTransferTx(t *testing.T) {
	transferType := "TOPUP"
	amount := int64(1000)
	srcBalance := int64(0)
	dstBalance := int64(0)
	n := 2
	errs := make(chan error)
	results := make(chan entity.Transfer)

	walletSrc := entity.Wallet{
		ID:           "001",
		AccountPhone: "+62000000001",
		Balance:      0,
		Type:         "ASSETS",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	walletDst := entity.Wallet{
		ID:           "002",
		AccountPhone: "+62000000002",
		Balance:      0,
		Type:         "LIABILITIES",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	for i := 0; i < n; i++ {
		fmt.Println("index: ", i)
		go func() {
			reference := fmt.Sprintf("%d", time.Now().UnixMilli())
			fmt.Println("reference: ", reference)
			transferId := uuid.New().String()

			now := time.Now()

			transfer := &entity.Transfer{
				ID:          transferId,
				SrcWalletID: walletSrc.ID,
				DstWalletID: walletDst.ID,
				Amount:      amount,
				Reference:   reference,
				Type:        transferType,
				CreatedAt:   now,
				UpdatedAt:   now,
			}

			entries := []entity.Entry{}

			srcEntry := entity.Entry{
				ID:            uuid.New().String(),
				WalletID:      walletSrc.ID,
				CreditAmount:  0,
				DebitAmount:   amount,
				BalanceBefore: srcBalance,
				BalanceAfter:  (srcBalance - amount),
				CorrelationID: reference,
				TransferID:    transferId,
				CreatedAt:     now,
				UpdatedAt:     now,
			}

			dstEntry := entity.Entry{
				ID:            uuid.New().String(),
				WalletID:      walletDst.ID,
				CreditAmount:  amount,
				DebitAmount:   0,
				BalanceBefore: dstBalance,
				BalanceAfter:  (dstBalance + amount),
				CorrelationID: reference,
				TransferID:    transferId,
				CreatedAt:     now,
				UpdatedAt:     now,
			}

			entries = append(entries, srcEntry, dstEntry)

			srcWalletUpdated := entity.WalletUpdateBalance{
				ID:        walletSrc.ID,
				Amount:    -amount,
				UpdatedAt: now,
			}

			dstWalletUpdated := entity.WalletUpdateBalance{
				ID:        walletDst.ID,
				Amount:    amount,
				UpdatedAt: now,
			}

			walletsToUpdate := []entity.WalletUpdateBalance{}
			walletsToUpdate = append(walletsToUpdate, srcWalletUpdated, dstWalletUpdated)

			updatedCounter := &entity.UpdateTransferCounter{
				WalletID:      dstWalletUpdated.ID,
				CountDaily:    1,
				AmountDaily:   amount,
				CountMonthly:  1,
				AmountMonthly: amount,
			}

			err := testStore.CreateTransferTx(
				context.Background(),
				transfer,
				entries,
				walletsToUpdate,
				updatedCounter,
			)

			errs <- err
			results <- *transfer

			srcBalance = srcBalance - amount
			dstBalance = dstBalance + amount

		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		assert.NoError(t, err)
	}
}
