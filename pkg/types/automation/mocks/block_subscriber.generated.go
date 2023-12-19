// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	automation "github.com/smartcontractkit/chainlink-common/pkg/types/automation"

	mock "github.com/stretchr/testify/mock"
)

// MockBlockSubscriber is an autogenerated mock type for the BlockSubscriber type
type MockBlockSubscriber struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *MockBlockSubscriber) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *MockBlockSubscriber) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Subscribe provides a mock function with given fields:
func (_m *MockBlockSubscriber) Subscribe() (int, chan automation.BlockHistory, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Subscribe")
	}

	var r0 int
	var r1 chan automation.BlockHistory
	var r2 error
	if rf, ok := ret.Get(0).(func() (int, chan automation.BlockHistory, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func() chan automation.BlockHistory); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(chan automation.BlockHistory)
		}
	}

	if rf, ok := ret.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Unsubscribe provides a mock function with given fields: _a0
func (_m *MockBlockSubscriber) Unsubscribe(_a0 int) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Unsubscribe")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockBlockSubscriber creates a new instance of MockBlockSubscriber. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBlockSubscriber(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBlockSubscriber {
	mock := &MockBlockSubscriber{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
