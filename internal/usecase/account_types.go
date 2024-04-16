package usecase

import (
	"errors"
	"fmt"

	"github.com/ahmadmilzam/ewallet/pkg/httpres"
	"github.com/ahmadmilzam/ewallet/pkg/validator"
)

type CreateAccountReqParams struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func (params *CreateAccountReqParams) Validate() (bool, error) {
	var err error

	if !validator.IsValidEmail(params.Email) {
		err = errors.New("CreateAccount: invalid amount params email")
		err = fmt.Errorf("%s: %w", httpres.InvalidAmount, err)
		return false, err
	}

	if !validator.IsValidPhone(params.Phone) {
		err = errors.New("CreateAccount: invalid req params phone")
		err = fmt.Errorf("%s: %w", httpres.InvalidPhone, err)
		return false, err
	}

	return true, nil
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
