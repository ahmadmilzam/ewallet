package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/ahmadmilzam/ewallet/config"
	"github.com/ahmadmilzam/ewallet/internal/entity"
	mockery "github.com/ahmadmilzam/ewallet/internal/store/_mock"
	httperrors "github.com/ahmadmilzam/ewallet/pkg/http-errors"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/mock"
)

// var loc, _ = time.LoadLocation("Asia/Jakarta")
// var err error
// loc, err = time.LoadLocation("Asia/Jakarta")
// var timestamp = time.Date(2024, time.Now().Month(), 1, 1, 1, 1, 1, loc)
// var timestamp = time.Date(2024, time.May, 6, 14, 2, 27, 0, time.Local)

func TestAppUsecase_CreateAccount(t *testing.T) {
	_ = config.Load("config", "../../")

	type seed struct {
		phone         string
		name          string
		email         string
		coaType       string
		role          string
		status        string
		createdAt     time.Time
		updatedAt     time.Time
		walletIDCash  string
		walletIDPoint string
	}

	var dataSeeder = seed{
		phone:         faker.E164PhoneNumber(),
		name:          faker.Name(),
		email:         faker.Email(),
		coaType:       AccountCOATypeLiabilities,
		role:          AccountRoleRegistered,
		status:        AccountStatusActive,
		createdAt:     time.Now(),
		updatedAt:     time.Now(),
		walletIDCash:  "wallet_id_cash",
		walletIDPoint: "wallet_id_point",
	}

	type args struct {
		ctx    context.Context
		params CreateAccountRequest
	}

	type mockResults struct {
		err error
	}

	tests := []struct {
		name        string
		callMock    bool
		args        args
		mockResults mockResults
		want        *CreateAccountResponse
	}{
		// TODO: Add test cases.
		{
			name: "Invalid params: missing {name}",
			args: args{
				ctx: context.Background(),
				params: CreateAccountRequest{
					Name:    "",
					Phone:   faker.E164PhoneNumber(),
					Email:   faker.Email(),
					COAType: AccountCOATypeLiabilities,
				},
			},
			mockResults: mockResults{
				err: errors.New("some error"),
			},
			want: &CreateAccountResponse{
				Success: false,
				Data:    nil,
				Error:   httperrors.GenerateError(httperrors.GenericBadRequest, "Params {name} is required"),
			},
		},
		{
			name:     "Account existed",
			callMock: true,
			args: args{
				ctx: context.Background(),
				params: CreateAccountRequest{
					Name:    "User existed",
					Phone:   dataSeeder.phone,
					Email:   "existed@domain.com",
					COAType: AccountCOATypeLiabilities,
				},
			},
			mockResults: mockResults{
				err: errors.New("sql: violates unique constraint"),
			},
			want: &CreateAccountResponse{
				Success: false,
				Data:    nil,
				Error:   httperrors.GenerateError(httperrors.DataDuplication, "Fail to create new account, account existed"),
			},
		},
		{
			name: "Fail to create new account due to unknown DB error",
			args: args{
				ctx: context.Background(),
				params: CreateAccountRequest{
					Name:    dataSeeder.name,
					Phone:   dataSeeder.phone,
					Email:   dataSeeder.email,
					COAType: AccountCOATypeLiabilities,
				},
			},
			mockResults: mockResults{
				err: errors.New("pq: some db error"),
			},
			want: &CreateAccountResponse{
				Success: false,
				Error:   httperrors.GenerateError(httperrors.GenericInternalError, "Fail to create new account"),
				Data:    nil,
			},
		},
		{
			name: "Successfully create new account",
			args: args{
				ctx: context.Background(),
				params: CreateAccountRequest{
					Name:    dataSeeder.name,
					Phone:   dataSeeder.phone,
					Email:   dataSeeder.email,
					COAType: dataSeeder.coaType,
				},
			},
			mockResults: mockResults{
				err: nil,
			},
			want: &CreateAccountResponse{
				Success: true,
				Error:   nil,
				Data: &AccountWalletData{
					Phone:     dataSeeder.phone,
					Name:      dataSeeder.name,
					Email:     dataSeeder.email,
					Role:      dataSeeder.role,
					Status:    dataSeeder.status,
					COAType:   dataSeeder.coaType,
					CreatedAt: dataSeeder.createdAt.Format(time.RFC3339),
					UpdatedAt: dataSeeder.updatedAt.Format(time.RFC3339),
					Wallets: []WalletSummary{
						{
							ID:      dataSeeder.walletIDCash,
							Type:    WalletTypeCash,
							Balance: int64(0),
						},
						{
							ID:      dataSeeder.walletIDPoint,
							Type:    WalletTypePoint,
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
				config: config.GetAppConfig(),
			}

			mockStore.On(
				"CreateAccountTx",
				tt.args.ctx,
				mock.AnythingOfType("*entity.Account"),
				mock.Anything, // will error if passing an slice of type here
				mock.AnythingOfType("*entity.TransferCounter"),
			).Return(tt.mockResults.err).Maybe()

			// got := u.CreateAccount(tt.args.ctx, tt.args.params)
			// assert.Equalf(t, tt.want, got, "%v must equal to %v", got, tt.want)

			if got := u.CreateAccount(tt.args.ctx, tt.args.params); !reflect.DeepEqual(got.Success, tt.want.Success) {
				t.Errorf("AppCreateAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppUsecase_GetAccount(t *testing.T) {
	_ = config.Load("config", "../../")
	appConfig := config.GetAppConfig()

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
							CreatedAt: time.Now(),
							UpdatedAt: time.Now(),
						},
						Wallet: entity.Wallet{
							ID:           "wallet_id",
							AccountPhone: "+6281200000000",
							Balance:      int64(0),
							Type:         WalletTypeCash,
							CreatedAt:    time.Now(),
							UpdatedAt:    time.Now(),
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
					CreatedAt: time.Now().Format(time.RFC3339),
					UpdatedAt: time.Now().Format(time.RFC3339),
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
				t.Errorf("AppGetAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
