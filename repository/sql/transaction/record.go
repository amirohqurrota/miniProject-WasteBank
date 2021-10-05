package transaction

import (
	"time"
	"wastebank-ca/bussines/transaction"
	_userRepo "wastebank-ca/repository/sql/users"

	//_depositRepo "wastebank-ca/repository/sql/wasteDeposit"
	"wastebank-ca/repository/sql/waste"

	"gorm.io/gorm"
)

type Transaction struct {
	ID int
	//gorm.Model
	UserID     int
	User       _userRepo.User `gorm:"foreignKey:UserID"`
	AdminID    int
	TypeID     int
	Type       TransactionType `gorm:"foreignKey:TypeID"`
	Date       time.Time
	DepositID  int
	Deposit    WasteDeposit `gorm:"foreignKey:DepositID"`
	TotalMoney int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type TransactionType struct {
	gorm.Model
	Name string
}

func toDomainType(rec *TransactionType) transaction.DomainType {
	return transaction.DomainType{
		ID:   int(rec.ID),
		Name: rec.Name,
	}
}

func fromDomainType(domain transaction.DomainType) TransactionType {
	return TransactionType{
		//ID:   domain.ID,
		Name: domain.Name,
	}
}

type WasteDeposit struct {
	ID          int
	WasteId     int
	Waste       waste.Waste `gorm:"foreignKey:WasteId"`
	TotalHeight int
	TotalMoney  int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func toDomain(rec *Transaction) transaction.DomainTransaction {
	return transaction.DomainTransaction{
		ID:         int(rec.ID),
		UserID:     rec.UserID,
		AdminID:    rec.AdminID,
		TypeID:     rec.TypeID,
		Date:       rec.Date,
		DepositID:  rec.DepositID,
		TotalMoney: rec.TotalMoney,
	}
}

func fromDomain(domain transaction.DomainTransaction) Transaction {
	return Transaction{
		ID:         domain.ID,
		UserID:     domain.UserID,
		AdminID:    domain.AdminID,
		TypeID:     domain.TypeID,
		Date:       domain.Date,
		DepositID:  domain.DepositID,
		TotalMoney: domain.TotalMoney,
	}
}
