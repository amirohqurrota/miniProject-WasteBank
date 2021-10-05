package transaction

type serviceTransaction struct {
	repository Repository
}

func NewService(repoTransaction Repository) Service {
	return &serviceTransaction{
		repository: repoTransaction,
	}
}

// func NewDeposit(depositData *DomainDeposit) (*DomainDeposit, error) {
// 	result, err := repository.NewDeposit()
// 	if err != nil {
// 		return &DomainDeposit{}, err
// 	}
// 	return result, nil
// }

func (servTransaction serviceTransaction) Append(transaction *DomainTransaction) (*DomainTransaction, error) {
	//response,err:=servTransaction.Append()
	//change total saldo user
	//addnewwaste deposit/debit
	result, err := servTransaction.repository.Insert(transaction)
	if err != nil {
		return &DomainTransaction{}, err
	}
	return result, nil
}

// func (servTransaction *serviceTransaction) Update(transaction *DomainTransaction) (*DomainTransaction, error) {
// 	//fmt.Println("id service", transaction.ID)
// 	result, err := servTransaction.repository.Update(transaction)
// 	if err != nil {
// 		return &DomainTransaction{}, err
// 	}
// 	return result, nil
// }

// func (servTransaction *serviceTransaction) FindByID(id int) (*DomainTransaction, error) {
// 	//fmt.Println("id service", id)
// 	result, err := servTransaction.repository.FindByID(id)
// 	if err != nil {
// 		return &DomainTransaction{}, err
// 	}
// 	return result, nil
// }
