package usecase

import (
	"context"
	"fmt"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/pkg/httpres"
)

type WalletUsecaseInterface interface {
	// CreateWallet(ctx context.Context, params CreateAccountReqParams) (entity.Account, entity.Wallet, error)
	GetWallets(ctx context.Context, p string) ([]WalletResBody, error)
}

// func (u *AppUsecase) CreateWallet(ctx context.Context, params CreateAccountReqParams) (entity.Account, entity.Wallet, error) {

// 	aID := uuid.New().String()
// 	wID := uuid.New().String()
// 	cAt := time.Now()
// 	uAt := time.Now()

// 	ac := entity.Account{
// 		ID:        aID,
// 		Name:      params.Name,
// 		Phone:     params.Phone,
// 		Email:     params.Email,
// 		Role:      "REGISTERED",
// 		Status:    "ACTIVE",
// 		CreatedAt: cAt,
// 		UpdatedAt: uAt,
// 	}

// 	wl := entity.Wallet{
// 		ID:        wID,
// 		AccountId: aID,
// 		Balance:   0.00,
// 		Type:      "CASH",
// 		CreatedAt: cAt,
// 		UpdatedAt: uAt,
// 	}

// 	err := u.store.CreateAccountTx(ctx, ac, wl)

// 	if err != nil {
// 		return entity.Account{}, entity.Wallet{}, err
// 	}

// 	return ac, wl, nil
// }

func (u *AppUsecase) GetWallets(ctx context.Context, p string) ([]WalletResBody, error) {
	wallets, err := u.store.FindWalletsByPhone(ctx, p)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", httpres.GenericInternalError, err)
	}

	if len(wallets) == 0 {
		err = fmt.Errorf("%s: GetWallet", httpres.GenericNotFound)
		return nil, err
	}

	return u.mapGetWalletsResponse(wallets), nil
}

func (u *AppUsecase) mapGetWalletsResponse(wallets []entity.Wallet) []WalletResBody {
	res := []WalletResBody{}

	for _, w := range wallets {
		res = append(res, WalletResBody{
			ID:           w.ID,
			AccountPhone: w.AccountPhone,
			Type:         w.Type,
			Balance:      w.Balance,
		})
	}

	return res
}
