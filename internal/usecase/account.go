package usecase

import (
	"context"
	"log/slog"
	"strings"
	"time"

	"github.com/ahmadmilzam/ewallet/internal/entity"
	httperrors "github.com/ahmadmilzam/ewallet/pkg/http-errors"
	"github.com/ahmadmilzam/ewallet/pkg/logger"
	"github.com/ahmadmilzam/ewallet/pkg/uuid"
)

type AccountUsecaseInterface interface {
	CreateAccount(ctx context.Context, params CreateAccountRequest) *CreateAccountResponse
	GetAccount(ctx context.Context, phone string) *GetAccountResponse
}

func (u *AppUsecase) CreateAccount(ctx context.Context, params CreateAccountRequest) *CreateAccountResponse {
	var err error

	validationErr := params.Validate()
	if validationErr != nil {
		msg := "Fail to create account"
		slog.Warn(msg, logger.ErrAttr(validationErr))
		return &CreateAccountResponse{
			Success: false,
			Error:   validationErr,
		}
	}

	createdAt := time.Now()
	updatedAt := time.Now()

	account := &entity.Account{
		Name:      params.Name,
		Phone:     params.Phone,
		Email:     params.Email,
		Role:      "REGISTERED",
		Status:    "ACTIVE",
		COAType:   params.COAType,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	wallets := []entity.Wallet{}

	walletCash := entity.Wallet{
		ID:           uuid.New().String(),
		AccountPhone: params.Phone,
		Balance:      0.00,
		Type:         "CASH",
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}

	walletPoint := entity.Wallet{
		ID:           uuid.New().String(),
		AccountPhone: params.Phone,
		Balance:      0.00,
		Type:         "POINT",
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}

	wallets = append(wallets, walletCash, walletPoint)

	counter := &entity.TransferCounter{
		WalletID:            walletCash.ID,
		CreditCountDaily:    0,
		CreditCountMonthly:  0,
		CreditAmountDaily:   0,
		CreditAmountMonthly: 0,
		CreatedAt:           createdAt,
		UpdatedAt:           updatedAt,
	}

	err = u.store.CreateAccountTx(ctx, account, wallets, counter)

	if err != nil {
		msg := "Fail to create new account, account existed"
		if strings.Contains(err.Error(), "violates unique constraint") {
			slog.WarnContext(ctx, msg, logger.ErrAttr(err))
			return &CreateAccountResponse{
				Success: false,
				Error:   httperrors.GenerateError(httperrors.DataDuplication, msg),
			}
		}
		msg = "Fail to create new account"
		slog.ErrorContext(ctx, msg, logger.ErrAttr(err))

		return &CreateAccountResponse{
			Success: false,
			Error:   httperrors.GenerateError(httperrors.GenericInternalError, msg),
		}
	}

	accountWallets := []entity.AccountWallet{}
	walletCashS := entity.Wallet{
		ID:      walletCash.ID,
		Type:    walletCash.Type,
		Balance: walletCash.Balance,
	}
	walletPointS := entity.Wallet{
		ID:      walletPoint.ID,
		Type:    walletPoint.Type,
		Balance: walletPoint.Balance,
	}

	accountWallets = append(accountWallets,
		entity.AccountWallet{
			Account: *account,
			Wallet:  walletCashS,
		}, entity.AccountWallet{
			Account: *account,
			Wallet:  walletPointS,
		},
	)

	data := u.mapCreateAccountWalletResponse(accountWallets)

	return &CreateAccountResponse{
		Success: true,
		Data:    data,
	}
}

func (u *AppUsecase) GetAccount(ctx context.Context, phone string) *GetAccountResponse {
	accountWallets, err := u.store.FindAccountAndWalletsById(ctx, phone)

	if err != nil {
		msg := "Fail to get account"
		slog.ErrorContext(ctx, msg, logger.ErrAttr(err))

		return &GetAccountResponse{
			Success: false,
			Error:   httperrors.GenerateError(httperrors.GenericInternalError, msg),
		}
	}

	if len(accountWallets) == 0 {
		msg := "Account not found"
		err := httperrors.GenerateError(httperrors.AccountNotFound, msg)
		slog.WarnContext(ctx, msg, logger.ErrAttr(err))

		return &GetAccountResponse{
			Success: false,
			Error:   err,
		}
	}

	data := u.mapCreateAccountWalletResponse(accountWallets)

	return &GetAccountResponse{
		Success: true,
		Data:    data,
	}
}

func (u *AppUsecase) mapCreateAccountWalletResponse(feeder []entity.AccountWallet) *AccountWalletData {
	res := &AccountWalletData{
		Phone:     feeder[0].Phone,
		Name:      feeder[0].Name,
		Email:     feeder[0].Email,
		Role:      feeder[0].Role,
		Status:    feeder[0].Status,
		COAType:   feeder[0].COAType,
		CreatedAt: feeder[0].Account.CreatedAt.Format(time.RFC3339),
		UpdatedAt: feeder[0].Account.UpdatedAt.Format(time.RFC3339),
		Wallets:   []WalletSummary{},
	}

	for _, v := range feeder {
		res.Wallets = append(res.Wallets, WalletSummary{
			ID:      v.ID,
			Type:    v.Type,
			Balance: v.Balance,
		})
	}

	return res
}
