package usecase

import (
	"github.com/ahmadmilzam/ewallet/pkg/array"
	httperrors "github.com/ahmadmilzam/ewallet/pkg/http-errors"
	"github.com/ahmadmilzam/ewallet/pkg/validator"
)

type CreateTransferRequest struct {
	SrcWallet    string `json:"src_account"`
	DstWallet    string `json:"dst_account"`
	Amount       int64  `json:"amount"`
	TransferType string `json:"transfer_type"`
	Reference    string `json:"reference,omitempty"`
}

func (params *CreateTransferRequest) Validate() *httperrors.Error {
	if params.DstWallet == params.SrcWallet {
		return httperrors.GenerateError(httperrors.TransferToSameAccount, "Can't transfer to same account")
	}

	if !validator.IsValidAmount(params.Amount) {
		return httperrors.GenerateError(httperrors.InvalidAmount, "Invalid params {amount}")
	}

	if !array.Contains(GetSupportedTransferType(), params.TransferType) {
		return httperrors.GenerateError(httperrors.InvalidTransferType, "Invalid params {transfer_type}")
	}

	return nil
}

type CreateTransferData struct {
	*CreateTransferRequest
	TransferID string   `json:"transfer_id"`
	CreatedAt  JSONTime `json:"created_at"`
}

type CreateTransferResponse struct {
	Success bool                `json:"success"`
	Error   *httperrors.Error   `json:"error,omitempty"`
	Data    *CreateTransferData `json:"data,omitempty"`
}
