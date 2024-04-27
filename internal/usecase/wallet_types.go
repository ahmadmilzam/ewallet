package usecase

type WalletResBody struct {
	ID           string `json:"id"`
	AccountPhone string `json:"account_phone"`
	Type         string `json:"type"`
	Balance      int64  `json:"balance"`
}
