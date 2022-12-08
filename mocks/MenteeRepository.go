// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	mentee "api-alta-dashboard/features/mentee"

	mock "github.com/stretchr/testify/mock"
)

// MenteeRepository is an autogenerated mock type for the RepositoryInterface type
type MenteeRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: input
func (_m *MenteeRepository) Create(input mentee.Core) (int, error) {
	ret := _m.Called(input)

	var r0 int
	if rf, ok := ret.Get(0).(func(mentee.Core) int); ok {
		r0 = rf(input)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(mentee.Core) error); ok {
		r1 = rf(input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: id
func (_m *MenteeRepository) Delete(id int) (int, error) {
	ret := _m.Called(id)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUser provides a mock function with given fields: email
func (_m *MenteeRepository) FindUser(email string) (mentee.Core, int, error) {
	ret := _m.Called(email)

	var r0 mentee.Core
	if rf, ok := ret.Get(0).(func(string) mentee.Core); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(mentee.Core)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(string) int); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(email)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetAll provides a mock function with given fields: queryStatus, queryIdClass, queryEdType
func (_m *MenteeRepository) GetAll(queryStatus string, queryIdClass string, queryEdType string) ([]mentee.Core, error) {
	ret := _m.Called(queryStatus, queryIdClass, queryEdType)

	var r0 []mentee.Core
	if rf, ok := ret.Get(0).(func(string, string, string) []mentee.Core); ok {
		r0 = rf(queryStatus, queryIdClass, queryEdType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]mentee.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(queryStatus, queryIdClass, queryEdType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllWithSearch provides a mock function with given fields: queryName, queryStatus, queryIdClass, queryEdType
func (_m *MenteeRepository) GetAllWithSearch(queryName string, queryStatus string, queryIdClass string, queryEdType string) ([]mentee.Core, error) {
	ret := _m.Called(queryName, queryStatus, queryIdClass, queryEdType)

	var r0 []mentee.Core
	if rf, ok := ret.Get(0).(func(string, string, string, string) []mentee.Core); ok {
		r0 = rf(queryName, queryStatus, queryIdClass, queryEdType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]mentee.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string) error); ok {
		r1 = rf(queryName, queryStatus, queryIdClass, queryEdType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: id
func (_m *MenteeRepository) GetById(id int) (mentee.Core, error) {
	ret := _m.Called(id)

	var r0 mentee.Core
	if rf, ok := ret.Get(0).(func(int) mentee.Core); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(mentee.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: input, id
func (_m *MenteeRepository) Update(input mentee.Core, id int) (int, error) {
	ret := _m.Called(input, id)

	var r0 int
	if rf, ok := ret.Get(0).(func(mentee.Core, int) int); ok {
		r0 = rf(input, id)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(mentee.Core, int) error); ok {
		r1 = rf(input, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMenteeRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMenteeRepository creates a new instance of MenteeRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMenteeRepository(t mockConstructorTestingTNewMenteeRepository) *MenteeRepository {
	mock := &MenteeRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
