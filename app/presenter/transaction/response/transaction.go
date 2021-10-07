package response

import (
	"time"
	"wastebank-ca/bussines/transactions"
)

type Deposit struct {
	ID          int `json:"id"`
	WasteId     int `json:"wasteId"`
	TotalHeight int `json:"totalHeight"`
}

func fromDomainDeposit(domain transactions.DomainDeposit) Deposit {
	return Deposit{
		ID:          domain.ID,
		WasteId:     domain.WasteId,
		TotalHeight: domain.TotalHeight,
	}
}

func fromDomainAllDeposit(domain []transactions.DomainDeposit) []Deposit {
	var result []Deposit
	for _, element := range domain {
		result = append(result, fromDomainDeposit(element))
	}
	return result
}

type Transaction struct {
	ID          int       `json:"id"`
	UserID      int       `json:"userId"`
	AdminID     int       `json:"adminId"`
	TypeID      int       `json:"typeId"`
	Date        time.Time `json:"date"`
	TotalMoney  int       `json:"totalMoney"`
	DepositID   int       `json:"depositID"`
	DataDeposit []Deposit `json:"dataDeposit"`
}

func FromDomainTrans(domain transactions.DomainTransaction) Transaction {
	if domain.TypeID == 1 {
		//fmt.Println(domain.DepositData[0])
		return Transaction{
			TypeID:      domain.TypeID,
			ID:          domain.ID,
			UserID:      domain.UserID,
			AdminID:     domain.AdminID,
			TotalMoney:  domain.TotalMoney,
			Date:        domain.Date,
			DataDeposit: fromDomainAllDeposit(domain.DepositData),
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

type TypeTransaction struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FromDomainType(domain transactions.DomainType) TypeTransaction {
	return TypeTransaction{
		ID:   domain.ID,
		Name: domain.Name,
	}
}
