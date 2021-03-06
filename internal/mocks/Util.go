// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	validator_v9 "gopkg.in/go-playground/validator.v9"
)

// Util is an autogenerated mock type for the Util type
type Util struct {
	mock.Mock
}

// Struct provides a mock function with given fields: payload
func (_m *Util) Struct(payload interface{}) error {
	ret := _m.Called(payload)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Validate provides a mock function with given fields:
func (_m *Util) Validate() *validator_v9.Validate {
	ret := _m.Called()

	var r0 *validator_v9.Validate
	if rf, ok := ret.Get(0).(func() *validator_v9.Validate); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*validator_v9.Validate)
		}
	}

	return r0
}
