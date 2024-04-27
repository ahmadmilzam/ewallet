package store

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/pkg/uuid"
)

func TestStoreTransfer_CreateTransferTx(t *testing.T) {
	transferType := "TOPUP"
	amount := int64(10000)
	srcBalance := int64(0)
	dstBalance := int64(0)
	n := 10
	errs := make(chan error)
	results := make(chan string)

	walletSrc := entity.Wallet{
		ID:           "001",
		AccountPhone: "+62123123123",
		Balance:      0,
		Type:         "ASSETS",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	_, dstWallets, _, _ := CreateAccount()

	for i := 0; i < n; i++ {
		go func() {
			reference := fmt.Sprintf("%d", time.Now().UnixMilli())
			fmt.Println("reference: ", reference)
			transferId := uuid.New().String()

			now := time.Now()

			transfer := &entity.Transfer{
				ID:          transferId,
				SrcWalletID: walletSrc.ID,
				DstWalletID: dstWallets[0].ID,
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
				WalletID:      dstWallets[0].ID,
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
				ID:        dstWallets[0].ID,
				Amount:    amount,
				UpdatedAt: now,
			}

			walletsToUpdate := []entity.WalletUpdateBalance{}
			walletsToUpdate = append(walletsToUpdate, srcWalletUpdated, dstWalletUpdated)

			updatedCounter := &entity.UpdateTransferCounter{
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

			if err != nil {
				errs <- err
			} else {
				results <- "ok"
			}

			srcBalance = srcBalance - amount
			dstBalance = dstBalance + amount

		}()
	}
}
