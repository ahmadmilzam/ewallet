package usecase

const (
	AccountRoleUnregistered = "UNREGISTERED"
	AccountRoleRegistered   = "REGISTERED"
	AccountRoleInternalCoa  = "INTERNAL_COA"

	AccountStatusPending  = "PENDING"
	AccountStatusActive   = "ACTIVE"
	AccountStatusInactive = "INACTIVE"
	AccountStatusBlocked  = "BLOCKED"

	AccountCOATypeAssets      = "ASSETS"
	AccountCOATypeExpenses    = "EXPENSES"
	AccountCOATypeLiabilities = "LIABILITIES"
	AccountCOATypeEquity      = "EQUITY"
	AccountCOATypeRevenue     = "REVENUE"

	WalletTypeCash  = "CASH"
	WalletTypePoint = "POINT"

	TransferTypeDefault    = "TRANSFER"
	TransferTypeTopup      = "TOPUP"
	TransferTypePayment    = "PAYMENT"
	TransferTypeFee        = "FEE"
	TransferTypeReversal   = "REVERSAL"
	TransferTypeCorrection = "CORRECTION"
)

func GetSupportedAccountRole() []string {
	return []string{
		AccountRoleUnregistered,
		AccountRoleRegistered,
		AccountRoleInternalCoa,
	}
}

func GetSupportedTransferType() []string {
	return []string{
		TransferTypeDefault,
		TransferTypeTopup,
		TransferTypePayment,
		TransferTypeReversal,
		TransferTypeCorrection,
	}
}

func GetSupportedAccountStatus() []string {
	return []string{
		AccountStatusPending,
		AccountStatusActive,
		AccountStatusInactive,
		AccountStatusBlocked,
	}
}

func GetSupportedAccountCOA() []string {
	return []string{
		AccountCOATypeAssets,
		AccountCOATypeExpenses,
		AccountCOATypeLiabilities,
		AccountCOATypeEquity,
		AccountCOATypeRevenue,
	}
}
