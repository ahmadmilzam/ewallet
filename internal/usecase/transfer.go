package usecase

import (
	"context"
)

type TransferUsecaseInterface interface {
	// CreateWallet(ctx context.Context, params CreateAccountReqParams) (entity.Account, entity.Wallet, error)
	CreateTransfer(ctx context.Context, params CreateTransferReqParams) (CreateTransferResBody, error)
}

func (u *AppUsecase) CreateTransfer(ctx context.Context, params *CreateTransferReqParams) (*CreateTransferResBody, error) {
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
