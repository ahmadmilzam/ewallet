package usecase

var (
	AccountRole = map[string]string{
		"registered":   "REGISTERED",
		"unregistered": "UNREGISTERED",
		"internal_coa": "INTERNAL_COA",
	}

	AccountStatus = map[string]string{
		"active":   "ACTIVE",
		"inactive": "INACTIVE",
		"blocked":  "BLOCKED",
	}

	WalletType = map[string]string{
		"cash":  "CASH",
		"point": "POINT",
	}

	JournalType = map[string]string{
		"transfer":   "NORMAL_TRANSFER",
		"reversal":   "REVERSAL_TRANSFER",
		"correction": "CORRECTION_TRANSFER",
	}
)

const (
	AccountRoleUnregistered = "UNREGISTERED"
	AccountRoleRegistered   = "REGISTERED"
	AccountRoleInternalCoa  = "INTERNAL_COA"

	AccountStatusActive   = "ACTIVE"
	AccountStatusInactive = "INACTIVE"
	AccountStatusBlocked  = "BLOCKED"

	WalletTypeCash  = "CASH"
	WalletTypePoint = "POINT"

	JournalTypeTransfer   = "NORMAL_TRANSFER"
	JournalTypeReversal   = "REVERSAL_TRANSFER"
	JournalTypeCorrection = "CORRECTION_TRANSFER"

	CurrencyIDR = "IDR"
)

func GetSupportedAccountRole() []string {
	return []string{
		AccountRoleUnregistered,
		AccountRoleRegistered,
		AccountRoleInternalCoa,
	}
}

func GetSupportedAccountStatus() []string {
	return []string{
		AccountStatusActive,
		AccountStatusInactive,
		AccountStatusBlocked,
	}
}
