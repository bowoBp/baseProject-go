// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"
)

// DbConnector is an autogenerated mock type for the DbConnector type
type DbConnector struct {
	mock.Mock
}

type DbConnector_Expecter struct {
	mock *mock.Mock
}

func (_m *DbConnector) EXPECT() *DbConnector_Expecter {
	return &DbConnector_Expecter{mock: &_m.Mock}
}

// ConnectDb provides a mock function with given fields: dialector
func (_m *DbConnector) ConnectDb(dialector gorm.Dialector) (*gorm.DB, error) {
	ret := _m.Called(dialector)

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(gorm.Dialector) *gorm.DB); ok {
		r0 = rf(dialector)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gorm.Dialector) error); ok {
		r1 = rf(dialector)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DbConnector_ConnectDb_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ConnectDb'
type DbConnector_ConnectDb_Call struct {
	*mock.Call
}

// ConnectDb is a helper method to define mock.On call
//  - dialector gorm.Dialector
func (_e *DbConnector_Expecter) ConnectDb(dialector interface{}) *DbConnector_ConnectDb_Call {
	return &DbConnector_ConnectDb_Call{Call: _e.mock.On("ConnectDb", dialector)}
}

func (_c *DbConnector_ConnectDb_Call) Run(run func(dialector gorm.Dialector)) *DbConnector_ConnectDb_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(gorm.Dialector))
	})
	return _c
}

func (_c *DbConnector_ConnectDb_Call) Return(_a0 *gorm.DB, _a1 error) *DbConnector_ConnectDb_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewDbConnector interface {
	mock.TestingT
	Cleanup(func())
}

// NewDbConnector creates a new instance of DbConnector. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDbConnector(t mockConstructorTestingTNewDbConnector) *DbConnector {
	mock := &DbConnector{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}