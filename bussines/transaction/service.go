package transaction

import (
	"fmt"
	_adminDomain "wastebank-ca/bussines/admin"
	_userDomain "wastebank-ca/bussines/users"
)

type serviceTransaction struct {
	repository  Repository
	userDomain  _userDomain.Service
	adminDomain _adminDomain.Service
}

func NewService(repoTransaction Repository, adminService _adminDomain.Service, userService _userDomain.Service) Service {
	return &serviceTransaction{
		repository:  repoTransaction,
		userDomain:  userService,
		adminDomain: adminService,
	}
}

func (servTransaction serviceTransaction) Append(transaction *DomainTransaction) (*DomainTransaction, error) {
	//update total saldo user
	fmt.Println("service id", transaction.UserID)
	_, updateError := servTransaction.userDomain.UpdateSaldo(transaction.UserID, transaction.TotalMoney)
	if updateError != nil {
		if updateError != nil {
			fmt.Println("service error update")
			return &DomainTransaction{}, updateError
		}
	}
	fmt.Println("service update aman")
	result, err := servTransaction.repository.Insert(transaction)
	fmt.Println(result.AdminID, "trans")
	if err != nil {
		return &DomainTransaction{}, err
	}
	return result, nil
}

func (servTransaction serviceTransaction) AddNewType(typeTransaction *DomainType) (*DomainType, error) {
	fmt.Println("in service")
	result, err := servTransaction.repository.AddNewType(typeTransaction)
	if err != nil {
		return &DomainType{}, err
	}
	return result, nil
}
