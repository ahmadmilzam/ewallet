package usecase

import (
	"fmt"
	"strings"

	"github.com/ahmadmilzam/ewallet/config"
	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/pkg/httpres"
)

type AppUsecaseInterface interface {
	AccountUsecaseInterface
	WalletUsecaseInterface
	TransferUsecaseInterface
}

type AppUsecase struct {
	store  entity.StoreQuerier
	config config.AppConfig
}

func NewAppUsecase(s entity.StoreQuerier, c config.AppConfig) AppUsecaseInterface {
	return &AppUsecase{
		store:  s,
		config: c,
	}
}
func (u *AppUsecase) wrapNotFoundErr(e error) error {
	isNotFound := strings.Contains(e.Error(), "no rows in result set")
	if isNotFound {
		return fmt.Errorf("%s: %w", httpres.GenericNotFound, e)
	}
	return fmt.Errorf("%s: %w", httpres.GenericInternalError, e)
}
