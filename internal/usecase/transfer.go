package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/ahmadmilzam/ewallet/pkg/array"
	"github.com/ahmadmilzam/ewallet/pkg/httpres"
)

type TransferUsecaseInterface interface {
	// CreateWallet(ctx context.Context, params CreateAccountReqParams) (entity.Account, entity.Wallet, error)
	CreateTransfer(ctx context.Context, params *CreateTransferReqParams) (*CreateTransferResBody, error)
}

func (u *AppUsecase) CreateTransfer(ctx context.Context, params *CreateTransferReqParams) (*CreateTransferResBody, error) {
	var (
		err           error
		paramsIsValid bool
		// srcAccount    = &entity.Account{}
		// dstAccount    = &entity.Account{}
		// srcWallet     = &entity.Wallet{}
		// dstWallet     = &entity.Wallet{}
	)

	// validate the req params
	paramsIsValid, err = validateReqParams(params)
	if !paramsIsValid && err != nil {
		return nil, err
	}

	// // get the account src and dst detail (with the wallets)
	// // check if the accounts exist, or return err
	// srcAccount, err = u.store.FindAccountById(ctx, params.SrcWallet)
	// if err != nil {
	// 	err = fmt.Errorf("%s: CreateTransfer: %w", httpres.AccountNotFound, err)
	// 	return nil, err
	// }

	// // get the account balance info
	// // TODO: !!! Utilize INTERNAL_COA account for topup & payment in next iteration
	// // calculate transfer amount and apply below limitation for account type != INTERNAL_COA
	// // -- ensure wallet src's balance not < 0 after transfer, or return error
	// // -- ensure wallet dst's balance not > threshold after transfer, or return error
	// srcWallet, err = u.store.FindWalletById(ctx, params.SrcWallet)
	// if err != nil {
	// 	err = fmt.Errorf("%s: CreateTransfer: %w", httpres.GenericInternalError, err)
	// 	return nil, err
	// }

	// dstAccount, err = u.store.FindAccountById(ctx, params.DstWallet)
	// if err != nil {
	// 	err = fmt.Errorf("%s: %w", httpres.AccountDstNotFound, err)
	// 	return nil, err
	// }

	// dstWallet, err = u.store.FindWalletById(ctx, params.SrcWallet)
	// if err != nil {
	// 	err = fmt.Errorf("%s: CreateTransfer: %w", httpres.GenericInternalError, err)
	// 	return nil, err
	// }

	// srcBalanceAferIsValid, srcBalanceAfter := validateBalanceAfter(srcAccount.Role, srcWallet)
	// dstBalanceAferIsValid, dstBalanceAfter := validateBalanceAfter(dstAccount.Role, dstWallet)

	// calculate the credit counter limit daily & monthly still below threshold
	// calculate the credit amount dlimit daily & monthly ensure still below threshold
	// check if reference exist, or create default value, e.g: {src}_{dst}_epochtime
	// generate default reference if not provided
	// generate correlation_id
	// make transfer db transaction
	// -- insert to journal
	// -- inserts to transfers
	// -- update transfer_counter for dst_account
	// remap transfer response based on above process (err/ok)
	return nil, nil
}

func validateReqParams(params *CreateTransferReqParams) (bool, error) {
	var err error

	if params.Amount < 0 {
		err = errors.New("CreateTransfer: invalid amount params")
		err = fmt.Errorf("%s: %w", httpres.InvalidAmount, err)
		return false, err
	}

	if !array.Contains(GetSupportedTransferType(), params.Type) {
		err = errors.New("CreateTransfer: invalid transfer type")
		err = fmt.Errorf("%s: %w", httpres.InvalidTransferType, err)
		return false, err
	}

	return true, nil
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

// func generateCorrelationId(max int) string {
// 	source := rand.NewSource(time.Now().UnixNano())
// 	r := rand.New(source)

// 	tNow := time.Now().UnixNano()

// 	random := r.Intn(max)
// 	return strconv.Itoa(int(tNow)) + strconv.Itoa(random)

// }
