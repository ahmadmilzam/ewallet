package usecase

import (
	"context"
)

type TransferUsecaseInterface interface {
	// CreateWallet(ctx context.Context, params CreateAccountReqParams) (entity.Account, entity.Wallet, error)
	CreateTransfer(ctx context.Context, params *CreateTransferReqParams) (*CreateTransferResBody, error)
}

func (u *AppUsecase) CreateTransfer(ctx context.Context, params *CreateTransferReqParams) (*CreateTransferResBody, error) {
	// validate the req params
	// get the account src and dst detail (with the wallets)
	// check if the accounts exist, or return err
	// get the account balance info
	// TODO: !!! Utilize INTERNAL_COA account for topup & payment in next iteration
	// calculate transfer amount and apply below limitation for account type != INTERNAL_COA
	// -- ensure wallet src's balance not < 0 after transfer, or return error
	// -- ensure wallet dst's balance not > threshold after transfer, or return error
	// calculate the credit counter limit daily & monthly still below threshold
	// calculate the credit amount dlimit daily & monthly ensure still below threshold
	// check if reference exist, or create default value, e.g: {src}_{dst}_epochtime
	// generate correlation_id
	// make transfer db transaction
	// -- insert to journal
	// -- inserts to transfers
	// -- update transfer_counter for dst_account
	// remap transfer response based on above process (err/ok)
	return nil, nil
}

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
