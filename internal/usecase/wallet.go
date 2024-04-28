package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/pkg/httpres"
)

type WalletUsecaseInterface interface {
	GetWallet(ctx context.Context, id string) (*WalletResBody, error)
	GetWallets(ctx context.Context, phone string) ([]WalletResBody, error)
}

func (u *AppUsecase) GetWallet(ctx context.Context, id string) (*WalletResBody, error) {
	wallet, err := u.store.FindWalletById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: GetWallet: %w", httpres.GenericNotFound, err)
		}

		return nil, fmt.Errorf("%s: GetWallet: FindWalletById: %w", httpres.GenericInternalError, err)
	}

	return &WalletResBody{
		ID:           wallet.ID,
		AccountPhone: wallet.AccountPhone,
		Type:         wallet.Type,
		Balance:      wallet.Balance,
	}, nil
}

func (u *AppUsecase) GetWallets(ctx context.Context, phone string) ([]WalletResBody, error) {
	wallets, err := u.store.FindWalletsByPhone(ctx, phone)
	if err != nil {
		return nil, fmt.Errorf("%s: FindWalletsByPhone: %w", httpres.GenericInternalError, err)
	}

	if len(wallets) == 0 {
		return nil, fmt.Errorf("%s: GetWallets: no results %s", httpres.GenericNotFound, phone)
	}

	return u.mapGetWalletsSuccessResponse(wallets), nil
}

func (u *AppUsecase) mapGetWalletsSuccessResponse(wallets []entity.Wallet) []WalletResBody {
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
