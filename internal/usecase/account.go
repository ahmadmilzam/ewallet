package usecase

type AccountPayload struct {
	Phone  string `json:"phone"`
	Role   string `json:"role"`
	Status string `json:"status"`
}

type Account struct {
	ID        string `json:"id"`
	Phone     string `json:"phone"`
	Role      string `json:"role"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// func CreateAccount(payload *AccountPayload) *Account {
// 	return
// }
