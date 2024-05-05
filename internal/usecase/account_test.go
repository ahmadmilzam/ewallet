package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/ahmadmilzam/ewallet/config"
	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/internal/store"
	mockery "github.com/ahmadmilzam/ewallet/internal/store/_mock"
	httperrors "github.com/ahmadmilzam/ewallet/pkg/http-errors"
	"github.com/stretchr/testify/mock"
)

func TestAppUsecase_CreateAccount(t *testing.T) {
	type fields struct {
		store  store.SQLStoreInterface
		config config.AppConfig
	}
	type args struct {
		ctx    context.Context
		params CreateAccountRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *CreateAccountResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &AppUsecase{
				store:  tt.fields.store,
				config: tt.fields.config,
			}
			if got := u.CreateAccount(tt.args.ctx, tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppUsecase.CreateAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppUsecase_GetAccount(t *testing.T) {
	_ = config.Load("config", "../../")
	appConfig := config.GetAppConfig()

	timestamp := time.Date(2024, time.Now().Month(), 1, 1, 1, 1, 1, time.Local)

	type args struct {
		ctx   context.Context
		phone string
	}

	type mockResults struct {
		accounts []entity.AccountWallet
		err      error
	}

	tests := []struct {
		name        string
		config      config.AppConfig
		mockMethod  string
		mockResults mockResults
		args        args
		want        *GetAccountResponse
	}{
		// TODO: Add test cases.
		{
			name:       "Fail to get account",
			config:     appConfig,
			mockMethod: "FindAccountAndWalletsById",
			mockResults: mockResults{
				accounts: []entity.AccountWallet{},
				err:      errors.New("error from db"),
			},
			args: args{
				ctx:   context.Background(),
				phone: "081200000000",
			},
			want: &GetAccountResponse{
				Success: false,
				Error: &httperrors.Error{
					Code:    httperrors.GenericInternalError,
					Message: "Fail to get account",
				},
				Data: nil,
			},
		},
		{
			name:       "Account not found",
			config:     appConfig,
			mockMethod: "FindAccountAndWalletsById",
			mockResults: mockResults{
				accounts: []entity.AccountWallet{},
				err:      nil,
			},
			args: args{
				ctx:   context.Background(),
				phone: "081200000000",
			},
			want: &GetAccountResponse{
				Success: false,
				Error: &httperrors.Error{
					Code:    httperrors.AccountNotFound,
					Message: "Account not found",
				},
				Data: nil,
			},
		},
		{
			name:       "Account found",
			config:     appConfig,
			mockMethod: "FindAccountAndWalletsById",
			mockResults: mockResults{
				accounts: []entity.AccountWallet{
					{
						Account: entity.Account{
							Phone:     "+6281200000000",
							Name:      "Random Name",
							Email:     "email@domain.com",
							Role:      AccountRoleRegistered,
							Status:    AccountStatusActive,
							COAType:   AccountCOATypeLiabilities,
							CreatedAt: timestamp,
							UpdatedAt: timestamp,
						},
						Wallet: entity.Wallet{
							ID:           "wallet_id",
							AccountPhone: "+6281200000000",
							Balance:      int64(0),
							Type:         WalletTypeCash,
							CreatedAt:    timestamp,
							UpdatedAt:    timestamp,
						},
					},
				},
				err: nil,
			},
			args: args{
				ctx:   context.Background(),
				phone: "+6281200000000",
			},
			want: &GetAccountResponse{
				Success: true,
				Error:   nil,
				Data: &AccountWalletData{
					Phone:     "+6281200000000",
					Name:      "Random Name",
					Email:     "email@domain.com",
					Role:      AccountRoleRegistered,
					Status:    AccountStatusActive,
					COAType:   AccountCOATypeLiabilities,
					CreatedAt: JSONTime(timestamp),
					UpdatedAt: JSONTime(timestamp),
					Wallets: []WalletSummary{
						{
							ID:      "wallet_id",
							Type:    WalletTypeCash,
							Balance: int64(0),
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStore := mockery.NewSQLStoreInterface(t)
			u := &AppUsecase{
				store:  mockStore,
				config: tt.config,
			}
			mockStore.On(tt.mockMethod, mock.Anything, tt.args.phone).Return(tt.mockResults.accounts, tt.mockResults.err)
			if got := u.GetAccount(tt.args.ctx, tt.args.phone); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppUsecase.GetAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppUsecase_mapCreateAccountWalletResponse(t *testing.T) {
	type fields struct {
		store  store.SQLStoreInterface
		config config.AppConfig
	}
	type args struct {
		feeder []entity.AccountWallet
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *AccountWalletData
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &AppUsecase{
				store:  tt.fields.store,
				config: tt.fields.config,
			}
			if got := u.mapCreateAccountWalletResponse(tt.args.feeder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppUsecase.mapCreateAccountWalletResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
