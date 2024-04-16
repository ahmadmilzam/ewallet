package usecase

import (
	"errors"
	"fmt"

	"github.com/ahmadmilzam/ewallet/pkg/array"
	"github.com/ahmadmilzam/ewallet/pkg/httpres"
	"github.com/ahmadmilzam/ewallet/pkg/validator"
)

type CreateTransferReqParams struct {
	SrcWallet string  `json:"src_account"`
	DstWallet string  `json:"dst_account"`
	Amount    float64 `json:"amount"`
	Type      string  `json:"type"`
	Reference string  `json:"reference,omitempty"`
}

func (params *CreateTransferReqParams) Validate() (bool, error) {
	var err error

	if !validator.IsValidAmount(params.Amount) {
		err = errors.New("CreateTransfer: invalid amount params")
		err = fmt.Errorf("%s: %w", httpres.InvalidAmount, err)
		return false, err
	}

	if !array.Contains(GetSupportedTransferType(), params.Type) {
		err = errors.New("CreateTransfer: invalid transfer type")
		err = fmt.Errorf("%s: %w", httpres.InvalidTransferType, err)
		return false, err
	}

	return true, nil
}

type CreateTransferResBody struct{}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	// Transfer    Transfer `json:"transfer"`
	// FromAccount Account  `json:"from_account"`
	// ToAccount   Account  `json:"to_account"`
	// FromEntry   Entry    `json:"from_entry"`
	// ToEntry     Entry    `json:"to_entry"`
}
