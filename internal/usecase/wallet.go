package usecase

import (
	"context"

	"github.com/ahmadmilzam/ewallet/internal/entity"
)

type WalletUsecaseInterface interface {
	// CreateWallet(ctx context.Context, params CreateAccountReqParams) (entity.Account, entity.Wallet, error)
	GetWallet(ctx context.Context, p string) ([]entity.Wallet, error)
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

func (u *AppUsecase) GetWallet(ctx context.Context, p string) ([]entity.Wallet, error) {
	var wErr error
	w, err := u.store.FindWalletsByPhone(ctx, p)

	if err != nil {
		wErr = u.wrapNotFoundErr(err)
		return nil, wErr
	}

	return w, nil
}
