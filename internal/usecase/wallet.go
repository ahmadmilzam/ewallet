package usecase

import (
	"context"
	"fmt"
	"strings"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/pkg/httpres"
)

type WalletUsecaseInterface interface {
	// CreateWallet(ctx context.Context, params CreateAccountReqParams) (entity.Account, entity.Wallet, error)
	GetWallet(ctx context.Context, id string) (*entity.Wallet, error)
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

func (u *AppUsecase) GetWallet(ctx context.Context, id string) (*entity.Wallet, error) {
	fmt.Println("usecase/wallet")
	w, err := u.store.FindWalletById(ctx, id)
	if err != nil {
		fmt.Println("error find wallet", err)
	}

	if err != nil && !strings.Contains(err.Error(), "no rows in result set") {
		err = fmt.Errorf("%s: %w", httpres.GenericInternalError, err)
		return nil, err
	}

	if err != nil {
		err = fmt.Errorf("%s: %w", httpres.GenericNotFound, err)
		return nil, err
	}

	return w, nil
}

// func generateCorrelationId(max int) string {
// 	source := rand.NewSource(time.Now().UnixNano())
// 	r := rand.New(source)

// 	tNow := time.Now().UnixNano()

// 	random := r.Intn(max)
// 	return strconv.Itoa(int(tNow)) + strconv.Itoa(random)

// }
