// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	validator "gopkg.in/go-playground/validator.v9"
)

// AuthorPostgresValidator is an autogenerated mock type for the AuthorPostgresValidator type
type AuthorPostgresValidator struct {
	mock.Mock
}

// AuthorMustExist provides a mock function with given fields: fl
func (_m *AuthorPostgresValidator) AuthorMustExist(fl validator.FieldLevel) bool {
	ret := _m.Called(fl)

	var r0 bool
	if rf, ok := ret.Get(0).(func(validator.FieldLevel) bool); ok {
		r0 = rf(fl)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
