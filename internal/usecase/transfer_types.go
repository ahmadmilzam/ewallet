package usecase

import (
	"errors"
	"fmt"

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

func (params *CreateTransferRequest) Validate() (bool, error) {
	var err error
	if params.DstWallet == params.SrcWallet {
		err = errors.New("CreateTransfer: cannot transfer to same account")
		err = fmt.Errorf("%s: %w", httperrors.TransferToSameAccount, err)
		return false, err
	}

	if !validator.IsValidAmount(params.Amount) {
		err = errors.New("CreateTransfer: invalid amount params")
		err = fmt.Errorf("%s: %w", httperrors.InvalidAmount, err)
		return false, err
	}

	if !array.Contains(GetSupportedTransferType(), params.TransferType) {
		err = errors.New("CreateTransfer: invalid transfer type")
		err = fmt.Errorf("%s: %w", httperrors.InvalidTransferType, err)
		return false, err
	}

	return true, nil
}

type CreateTransferData struct {
	*CreateTransferRequest
	TransferID string   `json:"transfer_id"`
	CreatedAt  JSONTime `json:"created_at"`
}

type CreateTransferResponse struct {
	Success bool                `json:"success"`
	Error   *httperrors.Error   `json:"error"`
	Data    *CreateTransferData `json:"data"`
}
