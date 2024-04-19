package usecase

type WalletResBody struct {
	ID           string  `json:"id"`
	AccountPhone string  `json:"account_phone"`
	Type         string  `json:"type"`
	Balance      float64 `json:"balance"`
}
