// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	newsApi "wastebank-ca/bussines/newsApi"

	mock "github.com/stretchr/testify/mock"

	transactions "wastebank-ca/bussines/transactions"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// AddNewType provides a mock function with given fields: typeTransaction
func (_m *Service) AddNewType(typeTransaction *transactions.DomainType) (*transactions.DomainType, error) {
	ret := _m.Called(typeTransaction)

	var r0 *transactions.DomainType
	if rf, ok := ret.Get(0).(func(*transactions.DomainType) *transactions.DomainType); ok {
		r0 = rf(typeTransaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*transactions.DomainType)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*transactions.DomainType) error); ok {
		r1 = rf(typeTransaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Append provides a mock function with given fields: transaction
func (_m *Service) Append(transaction *transactions.DomainTransaction) (*transactions.DomainTransaction, *newsApi.Domain, error) {
	ret := _m.Called(transaction)

	var r0 *transactions.DomainTransaction
	if rf, ok := ret.Get(0).(func(*transactions.DomainTransaction) *transactions.DomainTransaction); ok {
		r0 = rf(transaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*transactions.DomainTransaction)
		}
	}

	var r1 *newsApi.Domain
	if rf, ok := ret.Get(1).(func(*transactions.DomainTransaction) *newsApi.Domain); ok {
		r1 = rf(transaction)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*newsApi.Domain)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*transactions.DomainTransaction) error); ok {
		r2 = rf(transaction)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
