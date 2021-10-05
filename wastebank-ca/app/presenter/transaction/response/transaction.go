package response

import (
	"time"
	"wastebank-ca/bussines/transaction"
)

type Deposit struct {
	ID          int `json:"id"`
	WasteId     int `json:"wasteId"`
	TotalHeight int `json:"totalHeight"`
}

type Transaction struct {
	ID          int       `json:"id"`
	UserID      int       `json:"userId"`
	AdminID     int       `json:"adminId"`
	TypeID      int       `json:"typeId"`
	Date        time.Time `json:"date"`
	TotalMoney  int       `json:"totalMoney"`
	DepositID   int       `json:"depositID"`
	DataDeposit Deposit   `json:"dataDeposit"`
}

func FromDomain(domain transaction.DomainTransaction) Transaction {
	if domain.TypeID == 1 {
		return Transaction{
			TypeID:      domain.TypeID,
			ID:          domain.ID,
			UserID:      domain.UserID,
			AdminID:     domain.AdminID,
			TotalMoney:  domain.TotalMoney,
			Date:        domain.Date,
			DataDeposit: Deposit(domain.DepositData),
		}
	}
	return Transaction{
		TypeID:     domain.TypeID,
		ID:         domain.ID,
		UserID:     domain.UserID,
		AdminID:    domain.AdminID,
		TotalMoney: domain.TotalMoney,
		Date:       domain.Date,
	}
}
