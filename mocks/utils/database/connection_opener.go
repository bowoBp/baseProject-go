// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// ConnectionOpener is an autogenerated mock type for the ConnectionOpener type
type ConnectionOpener struct {
	mock.Mock
}

type ConnectionOpener_Expecter struct {
	mock *mock.Mock
}

func (_m *ConnectionOpener) EXPECT() *ConnectionOpener_Expecter {
	return &ConnectionOpener_Expecter{mock: &_m.Mock}
}

// OpenConnection provides a mock function with given fields: dialector, config
func (_m *ConnectionOpener) OpenConnection(dialector gorm.Dialector, config *gorm.Config) (*gorm.DB, error) {
	ret := _m.Called(dialector, config)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(gorm.Dialector, *gorm.Config) *gorm.DB); ok {
		r0 = rf(dialector, config)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gorm.Dialector, *gorm.Config) error); ok {
		r1 = rf(dialector, config)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConnectionOpener_OpenConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'OpenConnection'
type ConnectionOpener_OpenConnection_Call struct {
	*mock.Call
}

// OpenConnection is a helper method to define mock.On call
//  - dialector gorm.Dialector
//  - config *gorm.Config
func (_e *ConnectionOpener_Expecter) OpenConnection(dialector interface{}, config interface{}) *ConnectionOpener_OpenConnection_Call {
	return &ConnectionOpener_OpenConnection_Call{Call: _e.mock.On("OpenConnection", dialector, config)}
}

func (_c *ConnectionOpener_OpenConnection_Call) Run(run func(dialector gorm.Dialector, config *gorm.Config)) *ConnectionOpener_OpenConnection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(gorm.Dialector), args[1].(*gorm.Config))
	})
	return _c
}

func (_c *ConnectionOpener_OpenConnection_Call) Return(_a0 *gorm.DB, _a1 error) *ConnectionOpener_OpenConnection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewConnectionOpener interface {
	mock.TestingT
	Cleanup(func())
}

// NewConnectionOpener creates a new instance of ConnectionOpener. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConnectionOpener(t mockConstructorTestingTNewConnectionOpener) *ConnectionOpener {
	mock := &ConnectionOpener{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
