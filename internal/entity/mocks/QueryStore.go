// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/ahmadmilzam/ewallet/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// QueryStore is an autogenerated mock type for the QueryStore type
type QueryStore struct {
	mock.Mock
}

// CreateAccount provides a mock function with given fields: ctx, model
func (_m *QueryStore) CreateAccount(ctx context.Context, model entity.Account) (entity.Account, error) {
	ret := _m.Called(ctx, model)

	if len(ret) == 0 {
		panic("no return value specified for CreateAccount")
	}

	var r0 entity.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Account) (entity.Account, error)); ok {
		return rf(ctx, model)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.Account) entity.Account); ok {
		r0 = rf(ctx, model)
	} else {
		r0 = ret.Get(0).(entity.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.Account) error); ok {
		r1 = rf(ctx, model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateJournal provides a mock function with given fields: ctx, model
func (_m *QueryStore) CreateJournal(ctx context.Context, model entity.Journal) (entity.Journal, error) {
	ret := _m.Called(ctx, model)

	if len(ret) == 0 {
		panic("no return value specified for CreateJournal")
	}

	var r0 entity.Journal
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Journal) (entity.Journal, error)); ok {
		return rf(ctx, model)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.Journal) entity.Journal); ok {
		r0 = rf(ctx, model)
	} else {
		r0 = ret.Get(0).(entity.Journal)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.Journal) error); ok {
		r1 = rf(ctx, model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTransfer provides a mock function with given fields: ctx, model
func (_m *QueryStore) CreateTransfer(ctx context.Context, model entity.Transfer) (entity.Transfer, error) {
	ret := _m.Called(ctx, model)

	if len(ret) == 0 {
		panic("no return value specified for CreateTransfer")
	}

	var r0 entity.Transfer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Transfer) (entity.Transfer, error)); ok {
		return rf(ctx, model)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.Transfer) entity.Transfer); ok {
		r0 = rf(ctx, model)
	} else {
		r0 = ret.Get(0).(entity.Transfer)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.Transfer) error); ok {
		r1 = rf(ctx, model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateWallet provides a mock function with given fields: ctx, model
func (_m *QueryStore) CreateWallet(ctx context.Context, model entity.Wallet) (entity.Wallet, error) {
	ret := _m.Called(ctx, model)

	if len(ret) == 0 {
		panic("no return value specified for CreateWallet")
	}

	var r0 entity.Wallet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Wallet) (entity.Wallet, error)); ok {
		return rf(ctx, model)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.Wallet) entity.Wallet); ok {
		r0 = rf(ctx, model)
	} else {
		r0 = ret.Get(0).(entity.Wallet)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.Wallet) error); ok {
		r1 = rf(ctx, model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAccount provides a mock function with given fields: ctx, id
func (_m *QueryStore) DeleteAccount(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAccount")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteJournal provides a mock function with given fields: ctx, id
func (_m *QueryStore) DeleteJournal(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteJournal")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTransfer provides a mock function with given fields: ctx, id
func (_m *QueryStore) DeleteTransfer(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTransfer")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteWallet provides a mock function with given fields: ctx, id
func (_m *QueryStore) DeleteWallet(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteWallet")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAccountById provides a mock function with given fields: ctx, id
func (_m *QueryStore) FindAccountById(ctx context.Context, id string) (entity.Account, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindAccountById")
	}

	var r0 entity.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entity.Account, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.Account); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(entity.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAccountByPhone provides a mock function with given fields: ctx, phone
func (_m *QueryStore) FindAccountByPhone(ctx context.Context, phone string) (entity.Account, error) {
	ret := _m.Called(ctx, phone)

	if len(ret) == 0 {
		panic("no return value specified for FindAccountByPhone")
	}

	var r0 entity.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entity.Account, error)); ok {
		return rf(ctx, phone)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.Account); ok {
		r0 = rf(ctx, phone)
	} else {
		r0 = ret.Get(0).(entity.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, phone)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAccountWallets provides a mock function with given fields: ctx, wid
func (_m *QueryStore) FindAccountWallets(ctx context.Context, wid string) ([]entity.Wallet, error) {
	ret := _m.Called(ctx, wid)

	if len(ret) == 0 {
		panic("no return value specified for FindAccountWallets")
	}

	var r0 []entity.Wallet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]entity.Wallet, error)); ok {
		return rf(ctx, wid)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []entity.Wallet); ok {
		r0 = rf(ctx, wid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Wallet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, wid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAccounts provides a mock function with given fields: ctx
func (_m *QueryStore) FindAccounts(ctx context.Context) ([]entity.Account, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for FindAccounts")
	}

	var r0 []entity.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]entity.Account, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []entity.Account); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindJournal provides a mock function with given fields: ctx, id
func (_m *QueryStore) FindJournal(ctx context.Context, id string) (entity.Journal, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindJournal")
	}

	var r0 entity.Journal
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entity.Journal, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.Journal); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(entity.Journal)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindJournals provides a mock function with given fields: ctx
func (_m *QueryStore) FindJournals(ctx context.Context) ([]entity.Wallet, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for FindJournals")
	}

	var r0 []entity.Wallet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]entity.Wallet, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []entity.Wallet); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Wallet)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTransfer provides a mock function with given fields: ctx, id
func (_m *QueryStore) FindTransfer(ctx context.Context, id string) (entity.Transfer, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindTransfer")
	}

	var r0 entity.Transfer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entity.Transfer, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.Transfer); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(entity.Transfer)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindTransfers provides a mock function with given fields: ctx
func (_m *QueryStore) FindTransfers(ctx context.Context) ([]entity.Transfer, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for FindTransfers")
	}

	var r0 []entity.Transfer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]entity.Transfer, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []entity.Transfer); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]entity.Transfer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindWallet provides a mock function with given fields: ctx, id
func (_m *QueryStore) FindWallet(ctx context.Context, id string) (entity.Wallet, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindWallet")
	}

	var r0 entity.Wallet
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entity.Wallet, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.Wallet); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(entity.Wallet)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateJournal provides a mock function with given fields: ctx, model
func (_m *QueryStore) UpdateJournal(ctx context.Context, model entity.Journal) (entity.Journal, error) {
	ret := _m.Called(ctx, model)

	if len(ret) == 0 {
		panic("no return value specified for UpdateJournal")
	}

	var r0 entity.Journal
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Journal) (entity.Journal, error)); ok {
		return rf(ctx, model)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.Journal) entity.Journal); ok {
		r0 = rf(ctx, model)
	} else {
		r0 = ret.Get(0).(entity.Journal)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.Journal) error); ok {
		r1 = rf(ctx, model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTransfer provides a mock function with given fields: ctx, model
func (_m *QueryStore) UpdateTransfer(ctx context.Context, model entity.Transfer) (entity.Transfer, error) {
	ret := _m.Called(ctx, model)

	if len(ret) == 0 {
		panic("no return value specified for UpdateTransfer")
	}

	var r0 entity.Transfer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, entity.Transfer) (entity.Transfer, error)); ok {
		return rf(ctx, model)
	}
	if rf, ok := ret.Get(0).(func(context.Context, entity.Transfer) entity.Transfer); ok {
		r0 = rf(ctx, model)
	} else {
		r0 = ret.Get(0).(entity.Transfer)
	}

	if rf, ok := ret.Get(1).(func(context.Context, entity.Transfer) error); ok {
		r1 = rf(ctx, model)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpgradeAccount provides a mock function with given fields: ctx, id
func (_m *QueryStore) UpgradeAccount(ctx context.Context, id string) (entity.Account, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for UpgradeAccount")
	}

	var r0 entity.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (entity.Account, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) entity.Account); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(entity.Account)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewQueryStore creates a new instance of QueryStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQueryStore(t interface {
	mock.TestingT
	Cleanup(func())
}) *QueryStore {
	mock := &QueryStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}