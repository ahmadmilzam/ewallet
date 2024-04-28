package usecase

import (
	"github.com/ahmadmilzam/ewallet/config"
	"github.com/ahmadmilzam/ewallet/internal/entity"
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
