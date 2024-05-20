// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	types "github.com/cosmos/cosmos-sdk/types"
	mock "github.com/stretchr/testify/mock"
)

// AccountKeeper is an autogenerated mock type for the AccountKeeper type
type AccountKeeper struct {
	mock.Mock
}

// GetModuleAddress provides a mock function with given fields: moduleName
func (_m *AccountKeeper) GetModuleAddress(moduleName string) types.AccAddress {
	ret := _m.Called(moduleName)

	var r0 types.AccAddress
	if rf, ok := ret.Get(0).(func(string) types.AccAddress); ok {
		r0 = rf(moduleName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(types.AccAddress)
		}
	}

	return r0
}

// NewAccountKeeper creates a new instance of AccountKeeper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountKeeper(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountKeeper {
	mock := &AccountKeeper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
