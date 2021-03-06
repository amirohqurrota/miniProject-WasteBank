// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	newsApi "wastebank-ca/bussines/newsApi"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetNews provides a mock function with given fields:
func (_m *Repository) GetNews() (newsApi.Domain, error) {
	ret := _m.Called()

	var r0 newsApi.Domain
	if rf, ok := ret.Get(0).(func() newsApi.Domain); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(newsApi.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
