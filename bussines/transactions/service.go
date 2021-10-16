package transactions

import (
	_adminDomain "wastebank-ca/bussines/admin"
	_newsDomain "wastebank-ca/bussines/newsApi"
	_userDomain "wastebank-ca/bussines/users"
)

type serviceTransaction struct {
	repository  Repository
	userDomain  _userDomain.Service
	adminDomain _adminDomain.Service
	newsDomain  _newsDomain.Repository
}

func NewService(repoTransaction Repository, adminService _adminDomain.Service, userService _userDomain.Service, newsRepository _newsDomain.Repository) Service {
	return &serviceTransaction{
		repository:  repoTransaction,
		userDomain:  userService,
		adminDomain: adminService,
		newsDomain:  newsRepository,
	}
}

func (servTransaction serviceTransaction) Append(transaction *DomainTransaction) (*DomainTransaction, *_newsDomain.Domain, error) {
	//update total saldo user
	if transaction.TypeID == 2 {
		transaction.TotalMoney = -transaction.TotalMoney
	}
	_, updateUserError := servTransaction.userDomain.UpdateSaldo(transaction.UserID, transaction.TotalMoney)
	if updateUserError != nil {
		if updateUserError != nil {
			return &DomainTransaction{}, &_newsDomain.Domain{}, updateUserError
		}
	}

	//update total bonus admin
	// _, updateAdminError := servTransaction.adminDomain.UpdateBonus(transaction.UserID, transaction.TotalMoney)
	// if updateAdminError != nil {
	// 	if updateAdminError != nil {
	// 		return &DomainTransaction{}, &_newsDomain.Domain{}, updateAdminError
	// 	}
	// }

	//get News
	resultNews, err := servTransaction.newsDomain.GetNews()
	if err != nil {
		return &DomainTransaction{}, &_newsDomain.Domain{}, err
	}

	resultTrans, err := servTransaction.repository.Insert(transaction)
	if err != nil {
		return &DomainTransaction{}, &_newsDomain.Domain{}, err
	}
	return resultTrans, &resultNews, nil
}

func (servTransaction serviceTransaction) AddNewType(typeTransaction *DomainType) (*DomainType, error) {
	//fmt.Println("in service")
	result, err := servTransaction.repository.AddNewType(typeTransaction)
	if err != nil {
		return &DomainType{}, err
	}
	return result, nil
}
