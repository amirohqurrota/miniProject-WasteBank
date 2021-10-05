package response

import (
	"time"
	"wastebank-ca/bussines/transaction"
)

//aar (WasteDeposit,Transaction)

// type WasteDeposit struct {
// 	ID          int `json:"id"`
// 	WasteId     int `json:"wasteId"`
// 	TotalHeight int `json:"totalHeight"`
// }

type Transaction struct {
	ID         int       `json:"id"`
	UserID     int       `json:"userId"`
	AdminID    int       `json:"adminId"`
	TypeID     int       `json:"typeId"`
	Date       time.Time `json:"date"`
	TotalMoney int       `json:"totalMoney"`
	DepositID  int       `json:"depositID"`
	//DataDeposit WasteDeposit `json:"dataDeposit"`
}

// func FromDomainDeposit(domain transaction.DomainDeposit) WasteDeposit {
// 	return WasteDeposit{
// 		ID:          domain.ID,
// 		WasteId:     domain.WasteId,
// 		TotalHeight: domain.TotalHeight,
// 	}
// }

func FromDomain(domain transaction.DomainTransaction) Transaction {
	return Transaction{
		TypeID:     domain.TypeID,
		ID:         domain.ID,
		UserID:     domain.UserID,
		AdminID:    domain.AdminID,
		TotalMoney: domain.TotalMoney,
		Date:       domain.Date,
		//DataDeposit: WasteDeposit(domain.DepositData),
	}
}

// func FromDomain(domain transaction.DomainTransaction) Transaction {
// 	deposit:=FromDomainDeposit(domain.DepositData)
// 	if domain.TypeID == 1 {
// 		return Transaction{
// 			TypeID:      domain.TypeID,
// 			ID:          domain.ID,
// 			UserID:      domain.UserID,
// 			AdminID:     domain.AdminID,
// 			TotalMoney:  domain.TotalMoney,
// 			Date:        domain.Date,
// 			DataDeposit: deposit,
// 			}
// 		}
// 	}else {
// 		return Transaction{
// 			TypeID:     domain.TypeID,
// 			ID:         domain.ID,
// 			UserID:     domain.UserID,
// 			AdminID:    domain.AdminID,
// 			TotalMoney: domain.TotalMoney,
// 			Date:       domain.Date,
// 		}
// 	}
// }
