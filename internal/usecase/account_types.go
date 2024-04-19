package usecase

import (
	"errors"
	"fmt"

	"github.com/ahmadmilzam/ewallet/pkg/array"
	"github.com/ahmadmilzam/ewallet/pkg/httpres"
	"github.com/ahmadmilzam/ewallet/pkg/validator"
)

type CreateAccountReqParams struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	COAType string `json:"coa_type"`
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

	if !array.Contains(GetSupportedAccountCOA(), params.COAType) {
		err = errors.New("CreateAccount: invalid coa type")
		err = fmt.Errorf("%s: %w", httpres.InvalidCOAType, err)
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
	COAType   string          `json:"coa_type"`
	CreatedAt JSONTime        `json:"created_at"`
	UpdatedAt JSONTime        `json:"updated_at"`
	Wallets   []WalletSummary `json:"wallets"`
}

type WalletSummary struct {
	ID      string  `json:"id"`
	Type    string  `json:"type"`
	Balance float64 `json:"balance"`
}
