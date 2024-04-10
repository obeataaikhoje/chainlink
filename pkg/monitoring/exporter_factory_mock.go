// Code generated by mockery v2.38.0. DO NOT EDIT.

package monitoring

import mock "github.com/stretchr/testify/mock"

// ExporterFactoryMock is an autogenerated mock type for the ExporterFactory type
type ExporterFactoryMock struct {
	mock.Mock
}

// GetType provides a mock function with given fields:
func (_m *ExporterFactoryMock) GetType() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetType")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewExporter provides a mock function with given fields: _a0
func (_m *ExporterFactoryMock) NewExporter(_a0 Params) (Exporter, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for NewExporter")
	}

	var r0 Exporter
	var r1 error
	if rf, ok := ret.Get(0).(func(Params) (Exporter, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(Params) Exporter); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(Exporter)
		}
	}

	if rf, ok := ret.Get(1).(func(Params) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewExporterFactoryMock creates a new instance of ExporterFactoryMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewExporterFactoryMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *ExporterFactoryMock {
	mock := &ExporterFactoryMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
