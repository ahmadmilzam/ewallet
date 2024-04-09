package usecase

type CreateAccountReqParams struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type GetAccountReqParams struct {
	Phone string `uri:"phone"`
}
