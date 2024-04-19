package usecase

import (
	"errors"
	"fmt"

	"github.com/ahmadmilzam/ewallet/pkg/array"
	"github.com/ahmadmilzam/ewallet/pkg/httpres"
	"github.com/ahmadmilzam/ewallet/pkg/validator"
)

type TransferReqParams struct {
	SrcWallet    string  `json:"src_account,omitempty"`
	DstWallet    string  `json:"dst_account"`
	Amount       float64 `json:"amount"`
	TransferType string  `json:"transfer_type"`
	Reference    string  `json:"reference,omitempty"`
}

type TransferResBody struct {
	TransferReqParams
	TransferID string   `json:"transfer_id"`
	CreatedAt  JSONTime `json:"created_at"`
}

func (params *TransferReqParams) Validate() (bool, error) {
	var err error
	if params.DstWallet == params.SrcWallet {
		err = errors.New("CreateTransfer: cannot transfer to same account")
		err = fmt.Errorf("%s: %w", httpres.TransferToSameAccount, err)
		return false, err
	}
	if !validator.IsValidAmount(params.Amount) {
		err = errors.New("CreateTransfer: invalid amount params")
		err = fmt.Errorf("%s: %w", httpres.InvalidAmount, err)
		return false, err
	}

	if !array.Contains(GetSupportedTransferType(), params.TransferType) {
		err = errors.New("CreateTransfer: invalid transfer type")
		err = fmt.Errorf("%s: %w", httpres.InvalidTransferType, err)
		return false, err
	}

	return true, nil
}
