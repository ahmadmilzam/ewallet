package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/pkg/httpres"
	"github.com/ahmadmilzam/ewallet/pkg/uuid"
)

type AccountUsecaseInterface interface {
	CreateAccount(ctx context.Context, params CreateAccountReqParams) (*AccountWalletsResBody, error)
	GetAccount(ctx context.Context, phone string) (*AccountWalletsResBody, error)
}

func (u *AppUsecase) CreateAccount(ctx context.Context, p CreateAccountReqParams) (*AccountWalletsResBody, error) {
	createdAt := time.Now()
	updatedAt := time.Now()

	account := &entity.Account{
		Name:      p.Name,
		Phone:     p.Phone,
		Email:     p.Email,
		Role:      "REGISTERED",
		Status:    "ACTIVE",
		COAType:   p.COAType,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	wallets := []entity.Wallet{}

	walletCash := entity.Wallet{
		ID:           uuid.New().String(),
		AccountPhone: p.Phone,
		Balance:      0.00,
		Type:         "CASH",
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}

	walletPoint := entity.Wallet{
		ID:           uuid.New().String(),
		AccountPhone: p.Phone,
		Balance:      0.00,
		Type:         "POINT",
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}

	wallets = append(wallets, walletCash, walletPoint)

	counter := &entity.TransferCounter{
		WalletId:            walletCash.ID,
		CreditCountDaily:    0,
		CreditCountMonthly:  0,
		CreditAmountDaily:   0,
		CreditAmountMonthly: 0,
		CreatedAt:           createdAt,
		UpdatedAt:           updatedAt,
	}

	err := u.store.CreateAccountTx(ctx, account, wallets, counter)
	if err != nil {
		err = fmt.Errorf("%s: CreateAccount: %w", httpres.GenericInternalError, err)

		if strings.Contains(err.Error(), "duplicreatedAte key value violates unique constraint") {
			err = fmt.Errorf("%s: CreateAccount: account exists: %w", httpres.DataDuplication, err)
		}
		return nil, err
	}
	accountWallets := []entity.AccountWallet{}
	walletCashS := entity.Wallet{
		ID:      walletCash.ID,
		Type:    walletCash.Type,
		Balance: walletCash.Balance,
	}
	walletPointS := entity.Wallet{
		ID:      walletPoint.ID,
		Type:    walletPoint.Type,
		Balance: walletPoint.Balance,
	}

	accountWallets = append(accountWallets,
		entity.AccountWallet{
			Account: *account,
			Wallet:  walletCashS,
		}, entity.AccountWallet{
			Account: *account,
			Wallet:  walletPointS,
		},
	)

	response := u.mapCreateAccountWalletResponse(accountWallets)

	return response, nil
}

func (u *AppUsecase) GetAccount(ctx context.Context, phone string) (*AccountWalletsResBody, error) {
	accountWallets, err := u.store.FindAccountAndWalletsById(ctx, phone)

	if err != nil {
		return nil, fmt.Errorf("%s: GetAccount: %w", httpres.GenericInternalError, err)
	}

	if len(accountWallets) == 0 {
		err = errors.New("no rows in result set")
		return nil, fmt.Errorf("%s: GetAccount: %s : %w", httpres.GenericNotFound, phone, err)
	}

	response := u.mapCreateAccountWalletResponse(accountWallets)

	return response, nil
}

func (u *AppUsecase) mapCreateAccountWalletResponse(feeder []entity.AccountWallet) *AccountWalletsResBody {
	res := &AccountWalletsResBody{
		Phone:     feeder[0].Phone,
		Name:      feeder[0].Name,
		Email:     feeder[0].Email,
		Role:      feeder[0].Role,
		Status:    feeder[0].Status,
		COAType:   feeder[0].COAType,
		CreatedAt: JSONTime(feeder[0].Account.CreatedAt.Local()),
		UpdatedAt: JSONTime(feeder[0].Account.UpdatedAt.Local()),
		Wallets:   []WalletSummary{},
	}

	for _, v := range feeder {
		res.Wallets = append(res.Wallets, WalletSummary{
			ID:      v.ID,
			Type:    v.Type,
			Balance: v.Balance,
		})
	}

	return res
}
