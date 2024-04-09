package usecase

type CreateAccountReqParams struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type GetAccountReqParams struct {
	Phone string `uri:"phone"`
}
