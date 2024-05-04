package store_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/internal/usecase"
	"github.com/ahmadmilzam/ewallet/pkg/uuid"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTestExecTx_CreateTransferTx(t *testing.T) {
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
		ID:           "68878a7c-49b7-40be-a720-1c93c10e845b",
		AccountPhone: "+62812222222",
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

			updatedCounter := &entity.TransferCounter{
				WalletID:            dstWalletUpdated.ID,
				CreditCountDaily:    1,
				CreditAmountDaily:   amount,
				CreditCountMonthly:  1,
				CreditAmountMonthly: amount,
			}

			err := testStore.CreateTransferTx(
				context.Background(),
				transfer,
				entries,
				walletsToUpdate,
				updatedCounter,
				true,
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

func createAccount() (entity.Account, []entity.Wallet, entity.TransferCounter) {
	now := time.Now()
	account := entity.Account{
		Phone:     faker.E164PhoneNumber(),
		Name:      faker.FirstName(),
		Email:     faker.Email(),
		Role:      usecase.AccountRoleRegistered,
		Status:    usecase.AccountStatusActive,
		COAType:   usecase.AccountCOATypeLiabilities,
		CreatedAt: now,
		UpdatedAt: now,
	}

	wallets := createWalletForAccountt(account.Phone)
	counter := createTransferCounterr(wallets[0].ID)

	return account, wallets, counter
}

func createWalletForAccountt(phone string) []entity.Wallet {
	var wallets []entity.Wallet

	now := time.Now()
	cash := entity.Wallet{
		ID:           faker.UUIDDigit(),
		AccountPhone: phone,
		Balance:      0,
		Type:         usecase.WalletTypeCash,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	point := entity.Wallet{
		ID:           faker.UUIDDigit(),
		AccountPhone: phone,
		Balance:      0,
		Type:         usecase.WalletTypePoint,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	wallets = append(wallets, cash, point)
	return wallets
}

func createTransferCounterr(walletID string) entity.TransferCounter {
	return entity.TransferCounter{
		WalletID:            walletID,
		CreditCountDaily:    0,
		CreditCountMonthly:  0,
		CreditAmountDaily:   0,
		CreditAmountMonthly: 0,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}
}

func TestExecTx_CreateAccountTx(t *testing.T) {
	account, wallets, counter := createAccount()
	err1 := testStore.CreateAccountTx(context.Background(), &account, wallets, &counter)
	require.NoError(t, err1)
}
