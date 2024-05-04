package usecase

import (
	"github.com/ahmadmilzam/ewallet/pkg/array"
	httperrors "github.com/ahmadmilzam/ewallet/pkg/http-errors"
	"github.com/ahmadmilzam/ewallet/pkg/validator"
)

type CreateAccountRequest struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	COAType string `json:"coa_type"`
}

func (params *CreateAccountRequest) Validate() *httperrors.Error {
	if params.Name == "" {
		return httperrors.GenerateError(httperrors.GenericBadRequest, "Params {name} is required")
	}
	if params.Phone == "" {
		return httperrors.GenerateError(httperrors.GenericBadRequest, "Params {phone} is required")
	}
	if params.Email == "" {
		return httperrors.GenerateError(httperrors.GenericBadRequest, "Params {email} is required")
	}
	if params.COAType == "" {
		return httperrors.GenerateError(httperrors.GenericBadRequest, "Params {coa_type} is required")
	}

	if !validator.IsValidEmail(params.Email) {
		return httperrors.GenerateError(httperrors.InvalidEmail, "Invalid params {email}")
	}
	if !validator.IsValidPhone(params.Phone) {
		return httperrors.GenerateError(httperrors.InvalidPhone, "Invalid params {phone}")
	}
	if !array.Contains(GetSupportedAccountCOA(), params.COAType) {
		return httperrors.GenerateError(httperrors.InvalidCOAType, "Invalid params {coa_type}")
	}

	return nil
}

type AccountWalletData struct {
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

type CreateAccountResponse struct {
	Success bool               `json:"success"`
	Error   *httperrors.Error  `json:"error,omitempty"`
	Data    *AccountWalletData `json:"data,omitempty"`
}

type GetAccountResponse struct {
	Success bool               `json:"success"`
	Error   *httperrors.Error  `json:"error,omitempty"`
	Data    *AccountWalletData `json:"data,omitempty"`
}

type WalletSummary struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Balance int64  `json:"balance"`
}
