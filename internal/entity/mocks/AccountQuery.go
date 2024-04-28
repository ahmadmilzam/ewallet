// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/ahmadmilzam/ewallet/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// AccountQuery is an autogenerated mock type for the AccountQuery type
type AccountQuery struct {
	mock.Mock
}

// CreateAccount provides a mock function with given fields: ctx, account
func (_m *AccountQuery) CreateAccount(ctx context.Context, account *entity.Account) (*entity.Account, error) {
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
func (_m *AccountQuery) CreateAccountTx(ctx context.Context, account *entity.Account, wallets []entity.Wallet, counter *entity.TransferCounter) error {
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

// FindAccountAndWalletsById provides a mock function with given fields: ctx, id
func (_m *AccountQuery) FindAccountAndWalletsById(ctx context.Context, id string) ([]entity.AccountWallet, error) {
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
func (_m *AccountQuery) FindAccountById(ctx context.Context, id string) (*entity.Account, error) {
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
func (_m *AccountQuery) FindAccountForUpdateById(ctx context.Context, id string) (*entity.Account, error) {
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

// UpdateAccount provides a mock function with given fields: ctx, account
func (_m *AccountQuery) UpdateAccount(ctx context.Context, account *entity.Account) (*entity.Account, error) {
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

// NewAccountQuery creates a new instance of AccountQuery. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountQuery(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountQuery {
	mock := &AccountQuery{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}