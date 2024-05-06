package usecase

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/ahmadmilzam/ewallet/config"
	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/internal/store"
	mockery "github.com/ahmadmilzam/ewallet/internal/store/_mock"
	httperrors "github.com/ahmadmilzam/ewallet/pkg/http-errors"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
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
		walletIDCash:  "wallet_id_cash",
		walletIDPoint: "wallet_id_point",
	}

	type args struct {
		ctx    context.Context
		params CreateAccountRequest
	}

	type mockArgs struct {
		account *entity.Account
		wallets []entity.Wallet
		counter *entity.TransferCounter
	}

	type mockResults struct {
		err error
	}

	tests := []struct {
		name        string
		args        args
		mockMethod  string
		mockArgs    mockArgs
		mockResults mockResults
		want        *CreateAccountResponse
	}{
		// TODO: Add test cases.
		{
			name: "Invalid request params",
			args: args{
				ctx: context.Background(),
				params: CreateAccountRequest{
					Name:    "",
					Phone:   faker.E164PhoneNumber(),
					Email:   faker.Email(),
					COAType: AccountCOATypeLiabilities,
				},
			},
			mockMethod: "CreateAccountTx",
			mockArgs: mockArgs{
				account: &entity.Account{
					Phone:     "",
					Name:      "",
					Email:     "",
					Role:      "",
					Status:    "",
					COAType:   "",
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				},
				wallets: []entity.Wallet{
					{
						ID:           "wallet_id_cash",
						AccountPhone: dataSeeder.phone,
						Balance:      int64(0),
						Type:         WalletTypeCash,
						CreatedAt:    time.Now(),
						UpdatedAt:    time.Now(),
					},
					{
						ID:           "wallet_id_point",
						AccountPhone: dataSeeder.phone,
						Balance:      int64(0),
						Type:         WalletTypePoint,
						CreatedAt:    time.Now(),
						UpdatedAt:    time.Now(),
					},
				},
				counter: &entity.TransferCounter{
					WalletID:            "wallet_id_cash",
					CreditCountDaily:    int16(0),
					CreditCountMonthly:  int16(0),
					CreditAmountDaily:   int64(0),
					CreditAmountMonthly: int64(0),
					CreatedAt:           time.Now(),
					UpdatedAt:           time.Now(),
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
		// {
		// 	name: "Account existed",
		// 	args: args{
		// 		ctx: context.Background(),
		// 		params: CreateAccountRequest{
		// 			Name:    dataSeeder.name,
		// 			Phone:   dataSeeder.phone,
		// 			Email:   dataSeeder.email,
		// 			COAType: AccountCOATypeLiabilities,
		// 		},
		// 	},
		// 	mockMethod: "CreateAccountTx",
		// 	mockArgs: mockArgs{
		// 		account: &entity.Account{
		// 			Phone:     dataSeeder.phone,
		// 			Name:      dataSeeder.name,
		// 			Email:     dataSeeder.email,
		// 			Role:      dataSeeder.role,
		// 			Status:    dataSeeder.status,
		// 			COAType:   dataSeeder.coaType,
		// 			CreatedAt: time.Now(),
		// 			UpdatedAt: time.Now(),
		// 		},
		// 		wallets: []entity.Wallet{
		// 			{
		// 				ID:           "wallet_id_cash",
		// 				AccountPhone: dataSeeder.phone,
		// 				Balance:      int64(0),
		// 				Type:         WalletTypeCash,
		// 				CreatedAt:    time.Now(),
		// 				UpdatedAt:    time.Now(),
		// 			},
		// 			{
		// 				ID:           "wallet_id_point",
		// 				AccountPhone: dataSeeder.phone,
		// 				Balance:      int64(0),
		// 				Type:         WalletTypePoint,
		// 				CreatedAt:    time.Now(),
		// 				UpdatedAt:    time.Now(),
		// 			},
		// 		},
		// 		counter: &entity.TransferCounter{
		// 			WalletID:            "wallet_id_cash",
		// 			CreditCountDaily:    int16(0),
		// 			CreditCountMonthly:  int16(0),
		// 			CreditAmountDaily:   int64(0),
		// 			CreditAmountMonthly: int64(0),
		// 			CreatedAt:           time.Now(),
		// 			UpdatedAt:           time.Now(),
		// 		},
		// 	},
		// 	mockResults: mockResults{
		// 		err: errors.New("sql: violates unique constraint"),
		// 	},
		// 	want: &CreateAccountResponse{
		// 		Success: false,
		// 		Data:    nil,
		// 		Error:   httperrors.GenerateError(httperrors.DataDuplication, "Fail to create new account, account existed"),
		// 	},
		// },
		// {
		// 	name:   "Account created",
		// 	args: args{
		// 		ctx: context.Background(),
		// 		params: CreateAccountRequest{
		// 			Name:    dataSeeder.name,
		// 			Phone:   dataSeeder.phone,
		// 			Email:   dataSeeder.email,
		// 			COAType: AccountCOATypeLiabilities,
		// 		},
		// 	},
		// 	mockMethod: "CreateAccountTx",
		// 	mockArgs: mockArgs{
		// 		account: &entity.Account{
		// 			Phone:     dataSeeder.phone,
		// 			Name:      dataSeeder.name,
		// 			Email:     dataSeeder.email,
		// 			Role:      dataSeeder.role,
		// 			Status:    dataSeeder.status,
		// 			COAType:   dataSeeder.coaType,
		// 			CreatedAt: time.Now(),
		// 			UpdatedAt: time.Now(),
		// 		},
		// 		wallets: []entity.Wallet{
		// 			{
		// 				ID:           "wallet_id_cash",
		// 				AccountPhone: dataSeeder.phone,
		// 				Balance:      int64(0),
		// 				Type:         WalletTypeCash,
		// 				CreatedAt:    time.Now(),
		// 				UpdatedAt:    time.Now(),
		// 			},
		// 			{
		// 				ID:           "wallet_id_point",
		// 				AccountPhone: dataSeeder.phone,
		// 				Balance:      int64(0),
		// 				Type:         WalletTypePoint,
		// 				CreatedAt:    time.Now(),
		// 				UpdatedAt:    time.Now(),
		// 			},
		// 		},
		// 		counter: &entity.TransferCounter{
		// 			WalletID:            "wallet_id_cash",
		// 			CreditCountDaily:    int16(0),
		// 			CreditCountMonthly:  int16(0),
		// 			CreditAmountDaily:   int64(0),
		// 			CreditAmountMonthly: int64(0),
		// 			CreatedAt:           time.Now(),
		// 			UpdatedAt:           time.Now(),
		// 		},
		// 	},
		// 	mockResults: mockResults{
		// 		err: nil,
		// 	},
		// 	want: &CreateAccountResponse{
		// 		Success: true,
		// 		Data: &AccountWalletData{
		// 			Phone:     dataSeeder.phone,
		// 			Name:      dataSeeder.name,
		// 			Email:     dataSeeder.email,
		// 			Role:      dataSeeder.role,
		// 			Status:    dataSeeder.status,
		// 			COAType:   dataSeeder.coaType,
		// 			CreatedAt: time.Now().Format(time.RFC3339),
		// 			UpdatedAt: time.Now().Format(time.RFC3339),
		// 			Wallets: []WalletSummary{
		// 				{
		// 					ID:      dataSeeder.walletIDCash,
		// 					Type:    WalletTypeCash,
		// 					Balance: int64(0),
		// 				},
		// 				{
		// 					ID:      dataSeeder.walletIDPoint,
		// 					Type:    WalletTypePoint,
		// 					Balance: int64(0),
		// 				},
		// 			},
		// 		},
		// 		Error: nil,
		// 	},
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStore := mockery.NewSQLStoreInterface(t)
			u := &AppUsecase{
				store:  mockStore,
				config: config.GetAppConfig(),
			}

			if !strings.Contains(tt.name, "Invalid request params") {
				fmt.Println("tt name calling mock: ", tt.name)
				mockStore.On(tt.mockMethod, tt.args.ctx, tt.mockArgs.account, tt.mockArgs.wallets, tt.mockArgs.counter).Return(tt.mockResults.err)
			}

			got := u.CreateAccount(tt.args.ctx, tt.args.params)
			assert.Equal(t, tt.want, got)

			// if got := u.CreateAccount(tt.args.ctx, tt.args.params); !reflect.DeepEqual(got.Success, tt.want.Success) {
			// 	t.Errorf("AppCreateAccount() = %v, want %v", got, tt.want)
			// }
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
				t.Errorf("AppmapCreateAccountWalletResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
// receive
&entity.Account{
	Phone:"+710842617359",
	Name:"King Hailey Rolfson",
	Email:"EwVoOkE@fxvWkeS.net",
	Role:"REGISTERED",
	Status:"ACTIVE",
	COAType:"LIABILITIES",
	CreatedAt:time.Date(2024, time.May, 6, 9, 59, 58, 392043000, time.Local),
	UpdatedAt:time.Date(2024, time.May, 6, 9, 59, 58, 392044000, time.Local)}
// got
&entity.Account{
	Phone:"+710842617359",
	Name:"King Hailey Rolfson",
	Email:"EwVoOkE@fxvWkeS.net",
	Role:"REGISTERED",
	Status:"ACTIVE",
	COAType:"LIABILITIES",
	CreatedAt:time.Date(2024, time.May, 1, 1, 1, 1, 1, time.Local),
	UpdatedAt:time.Date(2024, time.May, 1, 1, 1, 1, 1, time.Local)
}
*/
