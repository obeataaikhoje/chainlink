// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	automation "github.com/smartcontractkit/chainlink-common/pkg/types/automation"
	mock "github.com/stretchr/testify/mock"
)

// MockEncoder is an autogenerated mock type for the Encoder type
type MockEncoder struct {
	mock.Mock
}

// Encode provides a mock function with given fields: _a0
func (_m *MockEncoder) Encode(_a0 ...automation.CheckResult) ([]byte, error) {
	_va := make([]interface{}, len(_a0))
	for _i := range _a0 {
		_va[_i] = _a0[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Encode")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(...automation.CheckResult) ([]byte, error)); ok {
		return rf(_a0...)
	}
	if rf, ok := ret.Get(0).(func(...automation.CheckResult) []byte); ok {
		r0 = rf(_a0...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(...automation.CheckResult) error); ok {
		r1 = rf(_a0...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Extract provides a mock function with given fields: _a0
func (_m *MockEncoder) Extract(_a0 []byte) ([]automation.ReportedUpkeep, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Extract")
	}

	var r0 []automation.ReportedUpkeep
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) ([]automation.ReportedUpkeep, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func([]byte) []automation.ReportedUpkeep); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]automation.ReportedUpkeep)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockEncoder creates a new instance of MockEncoder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEncoder(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEncoder {
	mock := &MockEncoder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
