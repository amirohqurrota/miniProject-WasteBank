// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	transactions "wastebank-ca/bussines/transactions"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// AddNewType provides a mock function with given fields: typeTransaction
func (_m *Repository) AddNewType(typeTransaction *transactions.DomainType) (*transactions.DomainType, error) {
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

// Insert provides a mock function with given fields: transaction
func (_m *Repository) Insert(transaction *transactions.DomainTransaction) (*transactions.DomainTransaction, error) {
	ret := _m.Called(transaction)

	var r0 *transactions.DomainTransaction
	if rf, ok := ret.Get(0).(func(*transactions.DomainTransaction) *transactions.DomainTransaction); ok {
		r0 = rf(transaction)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*transactions.DomainTransaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*transactions.DomainTransaction) error); ok {
		r1 = rf(transaction)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}