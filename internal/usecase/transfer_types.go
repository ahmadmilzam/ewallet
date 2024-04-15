package usecase

type CreateTransferReqParams struct {
	SrcWallet string  `json:"src_account"`
	DstWallet string  `json:"dst_account"`
	Amount    float64 `json:"amount"`
	Reference string  `json:"reference,omitempty"`
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
