package entity

//go:generate mockery --name StoreQuerier
type StoreQuerier interface {
	AccountQuery
	WalletQuery
	TransferCounterQuery
	TransferQuery
	EntryQuery
}
