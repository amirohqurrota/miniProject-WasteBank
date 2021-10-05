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
	ID        int
	UserID    int
	User      _userRepo.User `gorm:"foreignKey:UserID"`
	AdminID   int
	TypeID    int
	Type      TransactionType `gorm:"foreignKey:TypeID"`
	Date      time.Time
	DepositID int
	Deposit   WasteDeposit `gorm:"foreignKey:DepositID"`
	//DebitID    int
	//Debit      _debitRepo.Debit `gorm:"foreignKey:DebitID"`
	TotalMoney int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type TransactionType struct {
	gorm.Model
	Name string
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

// func (rec *Waste) tooDomain() waste.DomainWaste {
// 	return waste.DomainWaste{
// 		ID:            int(rec.ID),
// 		Name:          rec.Name,
// 		CategoryId:    int(rec.CategoryId),
// 		PurchasePrice: rec.PurchasePrice,
// 		TotalStock:    rec.TotalStock,
// 	}
// }

func fromDomain(domain transaction.DomainTransaction) Transaction {
	return Transaction{
		ID:        domain.ID,
		UserID:    domain.UserID,
		AdminID:   domain.AdminID,
		TypeID:    domain.TypeID,
		Date:      domain.Date,
		DepositID: domain.DepositID,
		//DebitID:    domain.DebitID,
		TotalMoney: domain.TotalMoney,
	}
}
