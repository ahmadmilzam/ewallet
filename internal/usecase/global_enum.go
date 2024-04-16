package usecase

const (
	AccountRoleUnregistered = "UNREGISTERED"
	AccountRoleRegistered   = "REGISTERED"
	AccountRoleInternalCoa  = "INTERNAL_COA"

	AccountStatusPending  = "PENDING"
	AccountStatusActive   = "ACTIVATED"
	AccountStatusInactive = "INACTIVATED"
	AccountStatusBlocked  = "BLOCKED"

	WalletTypeCash  = "CASH"
	WalletTypePoint = "POINT"

	JournalTypeTransfer   = "TRANSFER" // default
	JournalTypeTopup      = "TOPUP"
	JournalTypePayment    = "PAYMENT"
	JournalTypeReversal   = "REVERSAL"
	JournalTypeCorrection = "CORRECTION"

	CurrencyIDR = "IDR"
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
		AccountRoleUnregistered,
		AccountRoleRegistered,
		AccountRoleInternalCoa,
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
