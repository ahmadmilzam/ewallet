package usecase

import (
	"github.com/ahmadmilzam/ewallet/config"
	"github.com/ahmadmilzam/ewallet/internal/store"
)

//go:generate mockery --name "AppUsecaseInterface" --output "./_mock" --outpkg "mockery"
type AppUsecaseInterface interface {
	AccountUsecaseInterface
	WalletUsecaseInterface
	TransferUsecaseInterface
}

type AppUsecase struct {
	store  store.SQLStoreInterface
	config config.AppConfig
}

func NewAppUsecase(s store.SQLStoreInterface, c config.AppConfig) AppUsecaseInterface {
	return &AppUsecase{
		store:  s,
		config: c,
	}
}
