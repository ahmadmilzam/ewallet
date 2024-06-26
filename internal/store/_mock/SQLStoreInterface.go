// Code generated by mockery v2.42.1. DO NOT EDIT.

package mockery

import (
	context "context"

	entity "github.com/ahmadmilzam/ewallet/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// SQLStoreInterface is an autogenerated mock type for the SQLStoreInterface type
type SQLStoreInterface struct {
	mock.Mock
}

// CreateAccount provides a mock function with given fields: ctx, account
func (_m *SQLStoreInterface) CreateAccount(ctx context.Context, account *entity.Account) (*entity.Account, error) {
	ret := _m.Called(ctx, account)

	if len(ret) == 0 {
		panic("no return value specified for CreateAccount")
	}

	var r0 *entity.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Account) (*entity.Account, error)); ok {
		return rf(ctx, account)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Account) *entity.Account); ok {
		r0 = rf(ctx, account)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.Account) error); ok {
		r1 = rf(ctx, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateAccountTx provides a mock function with given fields: ctx, account, wallets, counter
func (_m *SQLStoreInterface) CreateAccountTx(ctx context.Context, account *entity.Account, wallets []entity.Wallet, counter *entity.TransferCounter) error {
	ret := _m.Called(ctx, account, wallets, counter)

	if len(ret) == 0 {
		panic("no return value specified for CreateAccountTx")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Account, []entity.Wallet, *entity.TransferCounter) error); ok {
		r0 = rf(ctx, account, wallets, counter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateCounter provides a mock function with given fields: ctx, counter
func (_m *SQLStoreInterface) CreateCounter(ctx context.Context, counter *entity.TransferCounter) (*entity.TransferCounter, error) {
	ret := _m.Called(ctx, counter)

	if len(ret) == 0 {
		panic("no return value specified for CreateCounter")
	}

	var r0 *entity.TransferCounter
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.TransferCounter) (*entity.TransferCounter, error)); ok {
		return rf(ctx, counter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.TransferCounter) *entity.TransferCounter); ok {
		r0 = rf(ctx, counter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.TransferCounter)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.TransferCounter) error); ok {
		r1 = rf(ctx, counter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateEntry provides a mock function with given fields: ctx, model
func (_m *SQLStoreInterface) CreateEntry(ctx context.Context, model *entity.Entry) (*entity.Entry, error) {
	ret := _m.Called(ctx, model)

	if len(ret) == 0 {
		panic("no return value specified for CreateEntry")
	}

	var r0 *entity.Entry
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Entry) (*entity.Entry, error)); ok {
		return rf(ctx, model)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Entry) *entity.Entry); ok {
		r0 = rf(ctx, model)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Entry)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.Entry) error); ok {
		r1 = rf(ctx, model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTransfer provides a mock function with given fields: ctx, transfer
func (_m *SQLStoreInterface) CreateTransfer(ctx context.Context, transfer *entity.Transfer) (*entity.Transfer, error) {
	ret := _m.Called(ctx, transfer)

	if len(ret) == 0 {
		panic("no return value specified for CreateTransfer")
	}

	var r0 *entity.Transfer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Transfer) (*entity.Transfer, error)); ok {
		return rf(ctx, transfer)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Transfer) *entity.Transfer); ok {
		r0 = rf(ctx, transfer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Transfer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.Transfer) error); ok {
		r1 = rf(ctx, transfer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTransferTx provides a mock function with given fields: ctx, transfer, entries, wallets, counter, lockCounter
func (_m *SQLStoreInterface) CreateTransferTx(ctx context.Context, transfer *entity.Transfer, entries []entity.Entry, wallets []entity.WalletUpdateBalance, counter *entity.TransferCounter, lockCounter bool) error {
	ret := _m.Called(ctx, transfer, entries, wallets, counter, lockCounter)

	if len(ret) == 0 {
		panic("no return value specified for CreateTransferTx")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Transfer, []entity.Entry, []entity.WalletUpdateBalance, *entity.TransferCounter, bool) error); ok {
		r0 = rf(ctx, transfer, entries, wallets, counter, lockCounter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateWallet provides a mock function with given fields: ctx, wallet
func (_m *SQLStoreInterface) CreateWallet(ctx context.Context, wallet *entity.Wallet) (*entity.Wallet, error) {
	ret := _m.Called(ctx, wallet)

	if len(ret) == 0 {
		panic("no return value specified for CreateWallet")
	}

	var r0 *entity.Wallet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Wallet) (*entity.Wallet, error)); ok {
		return rf(ctx, wallet)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Wallet) *entity.Wallet); ok {
		r0 = rf(ctx, wallet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Wallet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.Wallet) error); ok {
		r1 = rf(ctx, wallet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAccountAndWalletsById provides a mock function with given fields: ctx, id
func (_m *SQLStoreInterface) FindAccountAndWalletsById(ctx context.Context, id string) ([]entity.AccountWallet, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindAccountAndWalletsById")
	}

	var r0 []entity.AccountWallet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]entity.AccountWallet, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []entity.AccountWallet); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.AccountWallet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAccountById provides a mock function with given fields: ctx, id
func (_m *SQLStoreInterface) FindAccountById(ctx context.Context, id string) (*entity.Account, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindAccountById")
	}

	var r0 *entity.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Account, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Account); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAccountForUpdateById provides a mock function with given fields: ctx, id
func (_m *SQLStoreInterface) FindAccountForUpdateById(ctx context.Context, id string) (*entity.Account, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindAccountForUpdateById")
	}

	var r0 *entity.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Account, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Account); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindCounterById provides a mock function with given fields: ctx, id
func (_m *SQLStoreInterface) FindCounterById(ctx context.Context, id string) (*entity.TransferCounter, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindCounterById")
	}

	var r0 *entity.TransferCounter
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.TransferCounter, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.TransferCounter); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.TransferCounter)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindCounterForUpdateById provides a mock function with given fields: ctx, id
func (_m *SQLStoreInterface) FindCounterForUpdateById(ctx context.Context, id string) (*entity.TransferCounter, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindCounterForUpdateById")
	}

	var r0 *entity.TransferCounter
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.TransferCounter, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.TransferCounter); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.TransferCounter)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindEntryById provides a mock function with given fields: ctx, id
func (_m *SQLStoreInterface) FindEntryById(ctx context.Context, id string) (*entity.Entry, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindEntryById")
	}

	var r0 *entity.Entry
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Entry, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Entry); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Entry)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTransferById provides a mock function with given fields: ctx, id
func (_m *SQLStoreInterface) FindTransferById(ctx context.Context, id string) (*entity.Transfer, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindTransferById")
	}

	var r0 *entity.Transfer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Transfer, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Transfer); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Transfer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindWalletById provides a mock function with given fields: ctx, id
func (_m *SQLStoreInterface) FindWalletById(ctx context.Context, id string) (*entity.Wallet, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindWalletById")
	}

	var r0 *entity.Wallet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Wallet, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Wallet); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Wallet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindWalletForUpdateById provides a mock function with given fields: ctx, id
func (_m *SQLStoreInterface) FindWalletForUpdateById(ctx context.Context, id string) (*entity.Wallet, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindWalletForUpdateById")
	}

	var r0 *entity.Wallet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.Wallet, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.Wallet); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Wallet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindWalletsByPhone provides a mock function with given fields: ctx, phone
func (_m *SQLStoreInterface) FindWalletsByPhone(ctx context.Context, phone string) ([]entity.Wallet, error) {
	ret := _m.Called(ctx, phone)

	if len(ret) == 0 {
		panic("no return value specified for FindWalletsByPhone")
	}

	var r0 []entity.Wallet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]entity.Wallet, error)); ok {
		return rf(ctx, phone)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []entity.Wallet); ok {
		r0 = rf(ctx, phone)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Wallet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, phone)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAccount provides a mock function with given fields: ctx, account
func (_m *SQLStoreInterface) UpdateAccount(ctx context.Context, account *entity.Account) (*entity.Account, error) {
	ret := _m.Called(ctx, account)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAccount")
	}

	var r0 *entity.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Account) (*entity.Account, error)); ok {
		return rf(ctx, account)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Account) *entity.Account); ok {
		r0 = rf(ctx, account)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.Account) error); ok {
		r1 = rf(ctx, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCounter provides a mock function with given fields: ctx, counter
func (_m *SQLStoreInterface) UpdateCounter(ctx context.Context, counter *entity.TransferCounter) error {
	ret := _m.Called(ctx, counter)

	if len(ret) == 0 {
		panic("no return value specified for UpdateCounter")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.TransferCounter) error); ok {
		r0 = rf(ctx, counter)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateWallet provides a mock function with given fields: ctx, wallet
func (_m *SQLStoreInterface) UpdateWallet(ctx context.Context, wallet *entity.Wallet) error {
	ret := _m.Called(ctx, wallet)

	if len(ret) == 0 {
		panic("no return value specified for UpdateWallet")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Wallet) error); ok {
		r0 = rf(ctx, wallet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateWalletBalance provides a mock function with given fields: ctx, wallet
func (_m *SQLStoreInterface) UpdateWalletBalance(ctx context.Context, wallet *entity.WalletUpdateBalance) error {
	ret := _m.Called(ctx, wallet)

	if len(ret) == 0 {
		panic("no return value specified for UpdateWalletBalance")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.WalletUpdateBalance) error); ok {
		r0 = rf(ctx, wallet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSQLStoreInterface creates a new instance of SQLStoreInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSQLStoreInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *SQLStoreInterface {
	mock := &SQLStoreInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
