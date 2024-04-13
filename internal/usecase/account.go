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
	CreateAccount(ctx context.Context, params CreateAccountReqParams) (*entity.Account, *entity.Wallet, error)
	GetAccount(ctx context.Context, phone string) (*AccountWallets, error)
}

func (u *AppUsecase) CreateAccount(ctx context.Context, params CreateAccountReqParams) (*entity.Account, *entity.Wallet, error) {

	wID := uuid.New().String()
	cAt := time.Now()
	uAt := time.Now()

	ac := &entity.Account{
		Name:      params.Name,
		Phone:     params.Phone,
		Email:     params.Email,
		Role:      "REGISTERED",
		Status:    "ACTIVE",
		CreatedAt: cAt,
		UpdatedAt: uAt,
	}

	wl := &entity.Wallet{
		ID:           wID,
		AccountPhone: params.Phone,
		Balance:      0.00,
		Type:         "CASH",
		CreatedAt:    cAt,
		UpdatedAt:    uAt,
	}

	err := u.store.CreateAccountTx(ctx, ac, wl)

	if err != nil {
		return nil, nil, err
	}

	return ac, wl, nil
}

func (u *AppUsecase) GetAccount(ctx context.Context, p string) (*AccountWallets, error) {
	var err error
	var aac []entity.AccountWallet

	aac, err = u.store.FindAccountAndWalletsByPhone(ctx, p)
	switch {
	case err != nil && !strings.Contains(err.Error(), "no rows in result set"):
		err = fmt.Errorf("%s: %w", httpres.GenericInternalError, err)
	case err != nil:
		err = fmt.Errorf("%s: %w", httpres.GenericInternalError, err)
	default:
		err = nil
	}

	if err != nil {
		return nil, err
	}

	aws := &AccountWallets{}
	if err := formatFindAccountAndWalletsByPhone(aac, aws); err != nil {
		return nil, err
	}
	return aws, nil
}

func formatFindAccountAndWalletsByPhone(feeder []entity.AccountWallet, dest *AccountWallets) error {
	tz := feeder[0].CreatedAt.Local().Format(time.RFC3339)

	dest.Phone = feeder[0].Phone
	dest.Name = feeder[0].Name
	dest.Email = feeder[0].Email
	dest.Role = feeder[0].Role
	dest.Status = feeder[0].Status
	dest.CreatedAt = tz
	dest.Wallets = []WalletSummary{}

	for _, v := range feeder {
		dest.Wallets = append(dest.Wallets, WalletSummary{
			Type:    v.Type,
			Balance: v.Balance,
		})
	}
	return nil
}

// 2024-04-13T15:31:06.749112Z

// func generateCorrelationId(max int) string {
// 	source := rand.NewSource(time.Now().UnixNano())
// 	r := rand.New(source)

// 	tNow := time.Now().UnixNano()

// 	random := r.Intn(max)
// 	return strconv.Itoa(int(tNow)) + strconv.Itoa(random)

// }
