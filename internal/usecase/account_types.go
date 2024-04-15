package usecase

type CreateAccountReqParams struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type GetAccountReqParams struct {
	Phone string `uri:"phone"`
}

type AccountWalletsResBody struct {
	Phone     string          `json:"phone"`
	Name      string          `json:"name"`
	Email     string          `json:"email"`
	Role      string          `json:"role"`
	Status    string          `json:"status"`
	CreatedAt string          `json:"created_at"`
	Wallets   []WalletSummary `json:"wallets"`
}

type WalletSummary struct {
	Type    string  `json:"type"`
	Balance float64 `json:"balance"`
}
