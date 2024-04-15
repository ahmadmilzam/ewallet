package usecase

import (
	"context"
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
	cAt := time.Now()
	uAt := time.Now()

	a := &entity.Account{
		Name:      p.Name,
		Phone:     p.Phone,
		Email:     p.Email,
		Role:      "REGISTERED",
		Status:    "ACTIVE",
		CreatedAt: cAt,
		UpdatedAt: uAt,
	}

	ww := []entity.Wallet{}

	wc := entity.Wallet{
		ID:           uuid.New().String(),
		AccountPhone: p.Phone,
		Balance:      0.00,
		Type:         "CASH",
		CreatedAt:    cAt,
		UpdatedAt:    uAt,
	}

	wp := entity.Wallet{
		ID:           uuid.New().String(),
		AccountPhone: p.Phone,
		Balance:      0.00,
		Type:         "POINT",
		CreatedAt:    cAt,
		UpdatedAt:    uAt,
	}

	ww = append(ww, wc, wp)

	tc := &entity.TransferCounter{
		WalletId:            wc.ID,
		CountDaily:          0,
		CountMonthly:        0,
		CreditAmountDaily:   0,
		CreditAmountMonthly: 0,
		CreatedAt:           cAt,
		UpdatedAt:           uAt,
	}

	err := u.store.CreateAccountTx(ctx, a, ww, tc)
	if err != nil {
		err = fmt.Errorf("%s: CreateAccount: %w", httpres.GenericInternalError, err)

		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			err = fmt.Errorf("%s: CreateAccount: account exists: %w", httpres.DataDuplication, err)
		}
		return nil, err
	}

	awr := mapCreateAccountResponse(a, ww)

	return awr, nil
}

func (u *AppUsecase) GetAccount(ctx context.Context, p string) (*AccountWalletsResBody, error) {
	var wErr error

	aac, err := u.store.FindAccountAndWalletsById(ctx, p)

	if err != nil {
		wErr = u.wrapNotFoundErr(err)
		return nil, wErr
	}

	awr := formatFindAccountAndWalletsByPhone(aac)

	return awr, nil
}

func formatFindAccountAndWalletsByPhone(feeder []entity.AccountWallet) *AccountWalletsResBody {
	a := &entity.Account{
		Phone:     feeder[0].Phone,
		Name:      feeder[0].Name,
		Email:     feeder[0].Email,
		Role:      feeder[0].Role,
		Status:    feeder[0].Status,
		CreatedAt: feeder[0].CreatedAt,
		UpdatedAt: feeder[0].UpdatedAt,
	}

	ww := []entity.Wallet{}

	for _, v := range feeder {
		ww = append(ww, entity.Wallet{
			Type:    v.Type,
			Balance: v.Balance,
		})
	}

	res := mapCreateAccountResponse(a, ww)
	return res
}

func mapCreateAccountResponse(a *entity.Account, ww []entity.Wallet) *AccountWalletsResBody {

	tz := a.CreatedAt.Local().Format(time.RFC3339)
	res := &AccountWalletsResBody{}

	res.Phone = a.Phone
	res.Name = a.Name
	res.Email = a.Email
	res.Role = a.Role
	res.Status = a.Status
	res.CreatedAt = tz
	res.Wallets = []WalletSummary{}

	for _, v := range ww {
		res.Wallets = append(res.Wallets, WalletSummary{
			Type:    v.Type,
			Balance: v.Balance,
		})
	}

	return res
}
