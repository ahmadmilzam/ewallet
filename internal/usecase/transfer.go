package usecase

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/ahmadmilzam/ewallet/config"
	"github.com/ahmadmilzam/ewallet/internal/entity"
	httperrors "github.com/ahmadmilzam/ewallet/pkg/http-errors"
	"github.com/ahmadmilzam/ewallet/pkg/logger"
	"github.com/ahmadmilzam/ewallet/pkg/uuid"
)

type TransferUsecaseInterface interface {
	CreateTransfer(ctx context.Context, params CreateTransferRequest) *CreateTransferResponse
}

func (u *AppUsecase) CreateTransfer(ctx context.Context, params CreateTransferRequest) *CreateTransferResponse {
	// Validate request transfer params
	paramErr := params.Validate()
	if paramErr != nil {
		msg := "Fail to create account"
		slog.WarnContext(ctx, msg, logger.ErrAttr(paramErr))
		return &CreateTransferResponse{
			Success: false,
			Error:   paramErr,
		}
	}

	// Get the account src and dst detail (with the wallets)
	// and check if the accounts exist & ACTIVE, or return err
	srcAccount := u.GetAccount(ctx, params.SrcWallet)
	if !srcAccount.Success {
		return &CreateTransferResponse{
			Success: false,
			Error:   srcAccount.Error,
		}
	}

	if srcAccount.Data.Status != AccountStatusActive {
		msg := "Source account is inactive"
		err := httperrors.GenerateError(httperrors.InactiveAccount, msg)
		slog.Warn(msg, logger.ErrAttr(err))
		return &CreateTransferResponse{
			Success: false,
			Error:   err,
		}
	}
	srcCashWallet := srcAccount.Data.Wallets[0]

	// ensure that for topup transfer, the srcAccount type must be an assests
	if params.TransferType == TransferTypeTopup && srcAccount.Data.COAType != AccountCOATypeAssets {
		msg := fmt.Sprintf(
			"Account type must be an %s for operation %s",
			AccountCOATypeAssets,
			TransferTypeTopup,
		)
		err := httperrors.GenerateError(httperrors.InvalidCOAType, msg)
		slog.Warn(msg, logger.ErrAttr(err))
		return &CreateTransferResponse{
			Success: false,
			Error:   err,
		}
	}

	dstAccount := u.GetAccount(ctx, params.DstWallet)
	if !srcAccount.Success {
		return &CreateTransferResponse{
			Success: false,
			Error:   dstAccount.Error,
		}
	}

	if dstAccount.Data.Status != AccountStatusActive {
		msg := "Destination account is inactive"
		err := httperrors.GenerateError(httperrors.InactiveAccount, msg)
		slog.Warn(msg, logger.ErrAttr(err))
		return &CreateTransferResponse{
			Success: false,
			Error:   err,
		}
	}
	dstCashWallet := dstAccount.Data.Wallets[0]

	// defined the rules based on account's role
	rules := make(map[string]config.AccountConfig)
	rules["UNREGISTERED"] = u.config.Transfer.Unregistered
	rules["REGISTERED"] = u.config.Transfer.Registered
	creditRules := rules[dstAccount.Data.Role]

	// -- ensure wallet src's balance  >= 0 after transfer, or return error
	srcBalanceAfter := srcAccount.Data.Wallets[0].Balance - params.Amount
	if srcBalanceAfter < 0 && srcAccount.Data.Role != AccountRoleInternalCoa {
		msg := "insufficient balance"
		code := httperrors.InsufficientBalance
		slog.ErrorContext(ctx, fmt.Sprintf("%s: %s", code, msg))
		return &CreateTransferResponse{
			Success: false,
			Error:   httperrors.GenerateError(code, msg),
		}
	}

	// -- ensure wallet dst's balance not > threshold after transfer, or return error
	dstBalanceAfter := dstAccount.Data.Wallets[0].Balance + params.Amount
	if dstBalanceAfter > int64(creditRules.BalanceLimit) {
		msg := "destination account is exceeding balance limit"
		code := httperrors.ExceedBalanceAmount
		slog.ErrorContext(ctx, fmt.Sprintf("%s: %s", code, msg))
		return &CreateTransferResponse{
			Success: false,
			Error:   httperrors.GenerateError(code, msg),
		}
	}

	// -- ensure wallet credit count not > threshold
	// get and validate the counter, return err if exceeded the threshold
	transferCounter, err := u.store.FindCounterById(ctx, dstAccount.Data.Wallets[0].ID)
	if err != nil {
		code := httperrors.GenericInternalError
		msg := "Error finding counter"
		isNotFound := strings.Contains(err.Error(), "no rows in result set")
		if isNotFound {
			code = httperrors.CounterNotFound
			msg = "counter not found"
		}
		slog.ErrorContext(ctx, msg, logger.ErrAttr(err))

		return &CreateTransferResponse{
			Success: false,
			Error:   httperrors.GenerateError(code, msg),
		}
	}

	// calculate post transaction counter, will return updated counter data
	u.calculateCounter(params.Amount, transferCounter)

	// validate the credit counter limit daily & monthly still below threshold
	// validate the credit amount limit daily & monthly ensure still below threshold
	counterErr := u.validateCounter(transferCounter, creditRules)
	if counterErr != nil {
		msg := "Exceeding counter limit"
		slog.ErrorContext(ctx, msg, logger.ErrAttr(counterErr))
		return &CreateTransferResponse{
			Success: false,
			Error:   counterErr,
		}
	}

	// generate correlation_id
	correlationEntry := generateCorrelationId()

	// check if reference exist, or create default value, e.g: {src}_{dst}_epochtime
	// generate default reference if not provided
	if params.Reference == "" {
		params.Reference = fmt.Sprintf("%s_%d", correlationEntry, time.Now().UnixMilli())
	}

	// prepare transfer data
	transferId := uuid.New().String()
	now := time.Now()

	transfer := &entity.Transfer{
		ID:          transferId,
		SrcWalletID: srcCashWallet.ID,
		DstWalletID: dstCashWallet.ID,
		Amount:      params.Amount,
		Reference:   params.Reference,
		Type:        params.TransferType,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// generate entries
	entries := []entity.Entry{}

	srcEntry := entity.Entry{
		ID:            uuid.New().String(),
		WalletID:      srcCashWallet.ID,
		CreditAmount:  0,
		DebitAmount:   params.Amount,
		BalanceBefore: srcCashWallet.Balance,
		BalanceAfter:  srcBalanceAfter,
		CorrelationID: correlationEntry,
		TransferID:    transferId,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	dstEntry := entity.Entry{
		ID:            uuid.New().String(),
		WalletID:      dstCashWallet.ID,
		CreditAmount:  params.Amount,
		DebitAmount:   0,
		BalanceBefore: dstCashWallet.Balance,
		BalanceAfter:  dstBalanceAfter,
		CorrelationID: correlationEntry,
		TransferID:    transferId,
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	entries = append(entries, srcEntry, dstEntry)

	srcWalletUpdated := &entity.WalletUpdateBalance{
		ID:        srcCashWallet.ID,
		Amount:    -params.Amount,
		UpdatedAt: now,
	}

	dstWalletUpdated := &entity.WalletUpdateBalance{
		ID:        dstCashWallet.ID,
		Amount:    params.Amount,
		UpdatedAt: now,
	}

	wallets := []entity.WalletUpdateBalance{}
	wallets = append(wallets, *srcWalletUpdated, *dstWalletUpdated)

	// make transfer db transaction
	err = u.store.CreateTransferTx(
		ctx,
		transfer,
		entries,
		wallets,
		transferCounter,
		true,
	)

	if err != nil {
		msg := "Fail to create transfer"
		slog.ErrorContext(ctx, msg, logger.ErrAttr(err))
		return &CreateTransferResponse{
			Success: false,
			Error: &httperrors.Error{
				Code:    httperrors.GenericInternalError,
				Message: msg,
			},
		}
	}

	// return transfer response based on above process (err/ok)
	response := &CreateTransferResponse{
		Success: true,
		Data: &CreateTransferData{
			TransferID:            transferId,
			CreatedAt:             JSONTime(now),
			CreateTransferRequest: &params,
		},
	}
	return response
}

func (u *AppUsecase) calculateCounter(amount int64, counter *entity.TransferCounter) {
	currentDay := time.Now().Day()
	lastTransferDay := counter.UpdatedAt.Local().Day()

	currentMonth := time.Now().Month().String()
	lastTransferMonth := counter.UpdatedAt.Local().Month().String()

	// updateCounter := entity.UpdateTransferCounter{}
	// updateCounter.WalletID = counter.WalletId
	if currentDay == lastTransferDay {
		counter.CreditCountDaily++
		counter.CreditAmountDaily = counter.CreditAmountDaily + amount
		// updateCounter.CountDaily = 1
		// updateCounter.AmountDaily = amount
	} else {
		counter.CreditCountDaily = 1
		counter.CreditAmountDaily = amount
		// updateCounter.CountDaily = -(counter.CreditCountDaily - 1)
		// updateCounter.AmountDaily = -(counter.CreditAmountDaily - amount)
	}

	if currentMonth == lastTransferMonth {
		counter.CreditCountMonthly++
		counter.CreditAmountMonthly = counter.CreditAmountMonthly + amount
		// updateCounter.CountMonthly = 1
		// updateCounter.AmountMonthly = amount
	} else {
		counter.CreditCountMonthly = 1
		counter.CreditAmountMonthly = amount
		// updateCounter.CountDaily = -(counter.CreditCountMonthly - 1)
		// updateCounter.AmountDaily = -(counter.CreditAmountMonthly - amount)
	}
	now := time.Now()

	counter.UpdatedAt = now

	// updateCounter.UpdatedAt = now
}

func (u *AppUsecase) validateCounter(counter *entity.TransferCounter, rules config.AccountConfig) *httperrors.Error {
	if counter.CreditCountDaily > rules.CreditCountDailyLimit {
		return httperrors.GenerateError(httperrors.ExceedCountDaily, "exceeded credit count daily limit")
	}
	if counter.CreditCountMonthly > rules.CreditCountMonthlyLimit {
		return httperrors.GenerateError(httperrors.ExceedCountMonthly, "exceeded credit count monthly limit")
	}
	if counter.CreditAmountDaily > rules.CreditAmountDailyLimit {
		return httperrors.GenerateError(httperrors.ExceedAmountDaily, "exceeded credit amount daily limit")
	}
	if counter.CreditAmountMonthly > rules.CreditAmountMonthlyLimit {
		return httperrors.GenerateError(httperrors.ExceedAmountMonthly, "exceeded credit amount monthly limit")
	}
	return nil
}

func generateCorrelationId() string {
	const CORRELATION_MAX_DIGITS = 3
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	tNow := time.Now().UnixMilli()

	random := r.Intn(CORRELATION_MAX_DIGITS)
	return strconv.Itoa(int(tNow)) + strconv.Itoa(random)

}
