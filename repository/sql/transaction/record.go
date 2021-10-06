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
	ID int `gorm:"primaryKey"`
	//gorm.Model
	UserID      int
	User        _userRepo.User `gorm:"foreignKey:UserID"`
	AdminID     int
	TypeID      int
	Type        TransactionType `gorm:"foreignKey:TypeID"`
	Date        time.Time
	DepositID   int
	Deposit     WasteDeposit   `gorm:"foreignKey:DepositID"`
	DepositData []WasteDeposit `gorm:"foreignKey:ID"`
	TotalMoney  int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func toDomainTrans(rec *Transaction) transaction.DomainTransaction {
	if rec.TypeID == 1 {
		return transaction.DomainTransaction{
			ID:          int(rec.ID),
			UserID:      rec.UserID,
			AdminID:     rec.AdminID,
			TypeID:      rec.TypeID,
			Date:        rec.Date,
			DepositID:   rec.DepositID,
			DepositData: ToDomainAllDeposit(&rec.DepositData),
			TotalMoney:  rec.TotalMoney,
		}
	}
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

func fromDomainTrans(domain transaction.DomainTransaction) Transaction {
	return Transaction{
		ID:          domain.ID,
		UserID:      domain.UserID,
		AdminID:     domain.AdminID,
		TypeID:      domain.TypeID,
		Date:        domain.Date,
		DepositID:   domain.DepositID,
		TotalMoney:  domain.TotalMoney,
		DepositData: fromDomainAllDeposit(domain.DepositData),
	}
}

type TransactionType struct {
	gorm.Model
	Name string `gorm:"primaryKey"`
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
	ID          int `gorm:"primaryKey"`
	WasteId     int
	Waste       waste.Waste `gorm:"foreignKey:WasteId"`
	TotalHeight int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func ToDomainDeposit(rec *WasteDeposit) transaction.DomainDeposit {
	return transaction.DomainDeposit{
		ID:          int(rec.ID),
		WasteId:     rec.WasteId,
		TotalHeight: rec.TotalHeight,
	}
}

func ToDomainAllDeposit(rec *[]WasteDeposit) []transaction.DomainDeposit {
	var result []transaction.DomainDeposit
	for _, element := range *rec {
		result = append(result, ToDomainDeposit(&element))
	}
	return result
}

func fromDomainDeposit(domain transaction.DomainDeposit) WasteDeposit {
	return WasteDeposit{
		ID:          domain.ID,
		WasteId:     domain.WasteId,
		TotalHeight: domain.TotalHeight,
	}
}

func fromDomainAllDeposit(domain []transaction.DomainDeposit) []WasteDeposit {
	var result []WasteDeposit
	for _, element := range domain {
		result = append(result, fromDomainDeposit(element))
	}
	return result
}
