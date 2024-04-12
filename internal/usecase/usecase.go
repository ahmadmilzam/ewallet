package usecase

import "github.com/ahmadmilzam/ewallet/internal/entity"

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
