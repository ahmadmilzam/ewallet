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
	GetAccount(ctx context.Context, phone string) (*entity.Account, error)
}

func (u *AppUsecase) CreateAccount(ctx context.Context, params CreateAccountReqParams) (*entity.Account, *entity.Wallet, error) {

	aID := uuid.New().String()
	wID := uuid.New().String()
	cAt := time.Now()
	uAt := time.Now()

	ac := entity.Account{
		ID:        aID,
		Name:      params.Name,
		Phone:     params.Phone,
		Email:     params.Email,
		Role:      "REGISTERED",
		Status:    "ACTIVE",
		CreatedAt: cAt,
		UpdatedAt: uAt,
	}

	wl := entity.Wallet{
		ID:        wID,
		AccountId: aID,
		Balance:   0.00,
		Type:      "CASH",
		CreatedAt: cAt,
		UpdatedAt: uAt,
	}

	err := u.store.CreateAccountTx(ctx, ac, wl)

	if err != nil {
		return nil, nil, err
	}

	return &ac, &wl, nil
}

func (u *AppUsecase) GetAccount(ctx context.Context, phone string) (*entity.Account, error) {
	ac, err := u.store.FindAccountByPhone(ctx, phone)

	if err != nil && !strings.Contains(err.Error(), "no rows in result set") {
		err = fmt.Errorf("%s: %w", httpres.GenericInternalError, err)
		return nil, err
	}

	if err != nil {
		err = fmt.Errorf("%s: %w", httpres.GenericNotFound, err)
		return nil, err
	}

	return &ac, nil
}

// func generateCorrelationId(max int) string {
// 	source := rand.NewSource(time.Now().UnixNano())
// 	r := rand.New(source)

// 	tNow := time.Now().UnixNano()

// 	random := r.Intn(max)
// 	return strconv.Itoa(int(tNow)) + strconv.Itoa(random)

// }
