package entity

//go:generate mockery --name StoreQuerier
type StoreQueryInterface interface {
	AccountQuery
	WalletQuery
	TransferCounterQuery
	TransferQuery
	EntryQuery
}
