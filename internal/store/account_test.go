package store_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/pkg/randomizer"
	"github.com/ahmadmilzam/ewallet/pkg/uuid"
	"github.com/stretchr/testify/require"
)

func CreateAccount() (entity.Account, []entity.Wallet, entity.TransferCounter, error) {
	now := time.Now()
	account := entity.Account{}
	err := randomizer.GenerateAccountData(&account)
	if err != nil {
		return entity.Account{}, nil, entity.TransferCounter{}, err
	}
	account.Phone = fmt.Sprintf("+%s ", account.Phone)
	account.CreatedAt = now
	account.UpdatedAt = now

	wallets := createWalletForAccountt(account.Phone)
	counter := createTransferCounterr(wallets[0].ID)

	return account, wallets, counter, nil
}

func createWalletForAccountt(phone string) []entity.Wallet {
	var wallets []entity.Wallet

	now := time.Now()
	cash := entity.Wallet{
		ID:           uuid.New().String(),
		AccountPhone: phone,
		Balance:      0.00,
		Type:         "CASH",
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	point := entity.Wallet{
		ID:           uuid.New().String(),
		AccountPhone: phone,
		Balance:      0.00,
		Type:         "CASH",
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	wallets = append(wallets, cash, point)
	return wallets
}

func createTransferCounterr(walletID string) entity.TransferCounter {
	return entity.TransferCounter{
		WalletId:            walletID,
		CreditCountDaily:    0,
		CreditCountMonthly:  0,
		CreditAmountDaily:   0,
		CreditAmountMonthly: 0,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}
}

func TestStoreAccount_CreateAccountTx(t *testing.T) {
	account, wallets, counter, err := CreateAccount()
	require.NoError(t, err)

	err1 := testStore.CreateAccountTx(context.Background(), &account, wallets, &counter)

	require.NoError(t, err1)
}
