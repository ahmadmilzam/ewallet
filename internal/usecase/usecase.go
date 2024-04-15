package usecase

import (
	"fmt"
	"strings"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/pkg/httpres"
)

type AppUsecaseInterface interface {
	AccountUsecaseInterface
	WalletUsecaseInterface
}

type AppUsecase struct {
	store entity.StoreQuerier
}

func NewAppUsecase(store entity.StoreQuerier) AppUsecaseInterface {
	return &AppUsecase{
		store: store,
	}
}
func (u *AppUsecase) wrapNotFoundErr(e error) error {
	isNotFound := strings.Contains(e.Error(), "no rows in result set")
	if isNotFound {
		return fmt.Errorf("%s: %w", httpres.GenericNotFound, e)
	}
	return fmt.Errorf("%s: %w", httpres.GenericInternalError, e)
}
