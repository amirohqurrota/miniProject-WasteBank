package transaction

import (
	"time"
)

type DomainDeposit struct {
	ID          int
	WasteId     int
	TotalHeight int
}
type DomainType struct {
	ID   int
	Name string
}

type DomainTransaction struct {
	ID          int
	UserID      int
	AdminID     int
	TypeID      int
	Date        time.Time
	TotalMoney  int
	DepositID   int
	DepositData []DomainDeposit
}

type Service interface {
	Append(transaction *DomainTransaction) (*DomainTransaction, error)
	AddNewType(typeTransaction *DomainType) (*DomainType, error)
	//NewDeposit(deposit *DomainDeposit) (*DomainDeposit, error)
	//Update(transaction *DomainTransaction) (*DomainTransaction, error)
	//FindByID(id int) (*DomainTransaction, error)
}

type Repository interface {
	Insert(transaction *DomainTransaction) (*DomainTransaction, error)
	AddNewType(typeTransaction *DomainType) (*DomainType, error)
	//NewDeposit(deposit *DomainDeposit) (*DomainDeposit, error)
	//Update(transaction *DomainTransaction) (*DomainTransaction, error)
	//FindByID(id int) (*DomainTransaction, error)
}
