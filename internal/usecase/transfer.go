package usecase

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/ahmadmilzam/ewallet/config"
	"github.com/ahmadmilzam/ewallet/internal/entity"
	"github.com/ahmadmilzam/ewallet/pkg/httpres"
	"github.com/ahmadmilzam/ewallet/pkg/uuid"
)

type TransferUsecaseInterface interface {
	// CreateWallet(ctx context.Context, params CreateAccountReqParams) (entity.Account, entity.Wallet, error)
	CreateTransfer(ctx context.Context, params *TransferReqParams) (*TransferResBody, error)
}

func (u *AppUsecase) CreateTransfer(ctx context.Context, params *TransferReqParams) (*TransferResBody, error) {
	// get the account src and dst detail (with the wallets)
	// check if the accounts exist & ACTIVE, or return err
	srcAccount, err := u.findActiveAccount(ctx, params.SrcWallet)
	if err != nil {
		return nil, err
	}
	srcCashWallet := srcAccount.Wallets[0]

	// ensure that for topup transfer, the srcAccount type must be an assests
	if params.TransferType == TransferTypeTopup && srcAccount.COAType != AccountCOATypeAssets {
		err = errors.New("src wallet for topup transfer must be an assets")
		return nil, fmt.Errorf("%s: CreateTransfer: %w", httpres.IncorrectAccountType, err)
	}

	dstAccount, err := u.findActiveAccount(ctx, params.DstWallet)
	if err != nil {
		return nil, err
	}
	dstCashWallet := dstAccount.Wallets[0]

	// defined the rules based on account's role
	rules := make(map[string]config.AccountConfig)
	rules["UNREGISTERED"] = u.config.Transfer.Unregistered
	rules["REGISTERED"] = u.config.Transfer.Registered
	creditRules := rules[dstAccount.Role]

	// TODO: !!! Utilize INTERNAL_COA account for topup & payment in next iteration
	// Calculate transfer amount and apply below limitations for account role != INTERNAL_COA
	// -- ensure wallet src's balance not < 0 after transfer, or return error
	// -- ensure wallet dst's balance not > threshold after transfer, or return error
	srcBalanceAfter, dstBalanceAfter, err := u.isBalanceAfterAllowed(params.Amount, srcAccount, dstAccount, creditRules)
	if err != nil {
		return nil, err
	}

	// -- ensure wallet credit count not > threshold
	// get and validate the counter, return err if exceeded the threshold
	transferCounter, err := u.store.FindCounterById(ctx, dstAccount.Wallets[0].ID)
	if err != nil {
		isNotFound := strings.Contains(err.Error(), "no rows in result set")
		if isNotFound {
			return nil, fmt.Errorf("%s: %w", httpres.CounterNotFound, err)
		}
		return nil, fmt.Errorf("%s: %w", httpres.GenericInternalError, err)
	}

	// calculate post transaction counter, will modify current counter
	u.calculateCounter(params.Amount, transferCounter)

	// validate the credit counter limit daily & monthly still below threshold
	// validate the credit amount limit daily & monthly ensure still below threshold
	err = u.validateCounter(transferCounter, creditRules)
	if err != nil {
		println("validate counter err: ", err.Error())
		return nil, err
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

	srcWalletUpdated := &entity.Wallet{
		ID:           srcCashWallet.ID,
		AccountPhone: params.SrcWallet,
		Balance:      srcBalanceAfter,
		Type:         srcCashWallet.Type,
		CreatedAt:    time.Time(srcAccount.CreatedAt),
		UpdatedAt:    now,
	}

	dstWalletUpdated := &entity.Wallet{
		ID:           dstCashWallet.ID,
		AccountPhone: params.DstWallet,
		Balance:      dstBalanceAfter,
		Type:         dstCashWallet.Type,
		CreatedAt:    time.Time(srcAccount.CreatedAt),
		UpdatedAt:    now,
	}

	wallets := make(map[string]entity.Wallet)
	wallets["src"] = *srcWalletUpdated
	wallets["dst"] = *dstWalletUpdated

	needSrcWalletLock := false
	needDstWalletLock := false

	if srcAccount.Role != AccountRoleInternalCoa {
		needSrcWalletLock = true
	}

	if dstAccount.Role != AccountRoleInternalCoa {
		needDstWalletLock = true
	}

	// make transfer db transaction
	err = u.store.CreateTransferTx(
		ctx,
		transfer,
		entries,
		wallets,
		transferCounter,
		needSrcWalletLock,
		needDstWalletLock,
	)

	if err != nil {
		return nil, fmt.Errorf("%s: CreateTransfer: fail to create transfer: %w", httpres.GenericInternalError, err)
	}

	// remap transfer response based on above process (err/ok)
	response := &TransferResBody{
		TransferReqParams: *params,
		TransferID:        transferId,
		CreatedAt:         JSONTime(now),
	}
	return response, nil
}

func (u *AppUsecase) findActiveAccount(ctx context.Context, phone string) (*AccountWalletsResBody, error) {
	account, err := u.GetAccount(ctx, phone)
	if err != nil {
		return nil, err
	}

	if account.Status != AccountStatusActive {
		err = errors.New("findActiveAccount: not an active account")
		return nil, fmt.Errorf("%s: %w", httpres.InactiveAccount, err)
	}

	return account, nil
}

func (u *AppUsecase) isBalanceAfterAllowed(amount float64, srcAccount *AccountWalletsResBody, dstAccount *AccountWalletsResBody, creditRules config.AccountConfig) (float64, float64, error) {
	srcBalanceAfter := srcAccount.Wallets[0].Balance - amount
	if srcBalanceAfter-amount < 0 && srcAccount.Role != AccountRoleInternalCoa {
		err := errors.New("isBalanceAfterAllowed: insufficient balance")
		err = fmt.Errorf("%s: CreateTransfer: %w", httpres.InsufficientBalance, err)
		return 0, 0, err
	}

	dstBalanceAfter := dstAccount.Wallets[0].Balance + amount
	if dstBalanceAfter > float64(creditRules.BalanceLimit) {
		err := errors.New("isAdditionAllowed: exceed balance limit for the role")
		err = fmt.Errorf("%s: CreateTransfer: %w", httpres.ExceedBalanceAmount, err)
		return 0, 0, err
	}

	return srcBalanceAfter, dstBalanceAfter, nil
}

func (u *AppUsecase) calculateCounter(amount float64, counter *entity.TransferCounter) {
	currentMonth := time.Now().Month().String()
	currentDay := time.Now().Day()
	lastTransferMonth := counter.UpdatedAt.Local().Month().String()
	lastTransferDay := counter.UpdatedAt.Local().Day()

	if currentDay == lastTransferDay {
		counter.CreditCountDaily++
		counter.CreditAmountDaily = counter.CreditAmountDaily + amount
	} else {
		counter.CreditCountDaily = 1
		counter.CreditAmountDaily = amount
	}

	if currentMonth == lastTransferMonth {
		counter.CreditCountMonthly++
		counter.CreditAmountMonthly = counter.CreditAmountMonthly + amount
	} else {
		counter.CreditCountMonthly = 1
		counter.CreditAmountMonthly = amount
	}

	counter.UpdatedAt = time.Now()
}

func (u *AppUsecase) validateCounter(counter *entity.TransferCounter, rules config.AccountConfig) error {
	if counter.CreditCountDaily > rules.CreditCountDailyLimit {
		err := errors.New("exceeded credit count daily limit")
		return fmt.Errorf("%s: validateCounter: %w", httpres.ExceedCountDaily, err)
	}
	if counter.CreditCountMonthly > rules.CreditCountMonthlyLimit {
		err := errors.New("exceeded credit count monthly limit")
		return fmt.Errorf("%s: validateCounter: %w", httpres.ExceedCountMonthly, err)
	}
	if counter.CreditAmountDaily > rules.CreditAmountDailyLimit {
		err := errors.New("exceeded credit amount daily limit")
		return fmt.Errorf("%s: validateCounter: %w", httpres.ExceedAmountDaily, err)
	}
	if counter.CreditAmountMonthly > rules.CreditAmountMonthlyLimit {
		err := errors.New("exceeded credit amount monthly limit")
		return fmt.Errorf("%s: validateCounter: %w", httpres.ExceedAmountMonthly, err)
	}
	return nil
}

// func validateBalanceAfter(role string, wallet *entity.Wallet) (bool, float64) {
// 	rules := config.AppConfig.Tra
// 	return true, 0.00
// }

// func calculateCounter() {
// 	lastTransactionTimestamp = transactionTimestamp;
// 	DateFormat dateFormatDaily = new SimpleDateFormat(DATE_FORMAT_DAILY);
// 	DateFormat dateFormatMonthly = new SimpleDateFormat(DATE_FORMAT_MONTHLY);
// 	String transactionDaily = dateFormatDaily.format(transactionTimestamp);
// 	String transactionMonthly = dateFormatMonthly.format(transactionTimestamp);
// 	if (transactionDaily.equals(lastTransactionDaily)) {
// 		amountDaily = amountDaily.add(amount);
// 		countDaily++;
// 	} else {
// 		lastTransactionDaily = transactionDaily;
// 		amountDaily = amount;
// 		countDaily = 1;
// 	}
// 	if (transactionMonthly.equals(lastTransactionMonthly)) {
// 		amountMonthly = amountMonthly.add(amount);
// 		countMonthly++;
// 	} else {
// 		lastTransactionMonthly = transactionMonthly;
// 		amountMonthly = amount;
// 		countMonthly = 1;
// 	}
// }

func generateCorrelationId() string {
	const CORRELATION_MAX_DIGITS = 3
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	tNow := time.Now().UnixMilli()

	random := r.Intn(CORRELATION_MAX_DIGITS)
	return strconv.Itoa(int(tNow)) + strconv.Itoa(random)

}
