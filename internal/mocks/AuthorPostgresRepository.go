// Code generated by mockery 2.7.4. DO NOT EDIT.

package mocks

import (
	model "github.com/ciazhar/project-layout-rest-postgres/pkg/author/model"
	mock "github.com/stretchr/testify/mock"

	query "github.com/ciazhar/project-layout-rest-postgres/third_party/query"
)

// AuthorPostgresRepository is an autogenerated mock type for the AuthorPostgresRepository type
type AuthorPostgresRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *AuthorPostgresRepository) Delete(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: q
func (_m *AuthorPostgresRepository) Fetch(q *query.Query) ([]model.Author, error) {
	ret := _m.Called(q)

	var r0 []model.Author
	if rf, ok := ret.Get(0).(func(*query.Query) []model.Author); ok {
		r0 = rf(q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Author)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*query.Query) error); ok {
		r1 = rf(q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *AuthorPostgresRepository) GetByID(id string) (model.Author, error) {
	ret := _m.Called(id)

	var r0 model.Author
	if rf, ok := ret.Get(0).(func(string) model.Author); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.Author)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: req
func (_m *AuthorPostgresRepository) Store(req *model.Author) error {
	ret := _m.Called(req)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Author) error); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: req
func (_m *AuthorPostgresRepository) Update(req *model.Author) error {
	ret := _m.Called(req)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Author) error); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}