// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	admin "wastebank-ca/bussines/admin"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetData provides a mock function with given fields: id, firstName, lastName, username
func (_m *Repository) GetData(id int, firstName string, lastName string, username string) (*admin.Domain, error) {
	ret := _m.Called(id, firstName, lastName, username)

	var r0 *admin.Domain
	if rf, ok := ret.Get(0).(func(int, string, string, string) *admin.Domain); ok {
		r0 = rf(id, firstName, lastName, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string, string, string) error); ok {
		r1 = rf(id, firstName, lastName, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: _a0
func (_m *Repository) Insert(_a0 *admin.Domain) (*admin.Domain, error) {
	ret := _m.Called(_a0)

	var r0 *admin.Domain
	if rf, ok := ret.Get(0).(func(*admin.Domain) *admin.Domain); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*admin.Domain) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0
func (_m *Repository) Update(_a0 *admin.Domain) (*admin.Domain, error) {
	ret := _m.Called(_a0)

	var r0 *admin.Domain
	if rf, ok := ret.Get(0).(func(*admin.Domain) *admin.Domain); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*admin.Domain) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBonus provides a mock function with given fields: id, bonus
func (_m *Repository) UpdateBonus(id int, bonus int) (*admin.Domain, error) {
	ret := _m.Called(id, bonus)

	var r0 *admin.Domain
	if rf, ok := ret.Get(0).(func(int, int) *admin.Domain); ok {
		r0 = rf(id, bonus)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(id, bonus)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
