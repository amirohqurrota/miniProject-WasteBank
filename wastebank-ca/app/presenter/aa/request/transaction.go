package request

import (
	"time"
	//_depositReq "wastebank-ca/app/presenter/wasteDeposit/request"
	"wastebank-ca/bussines/transaction"
)

type WasteDeposit struct {
	ID          int `json:"id"`
	WasteId     int `json:"wasteId"`
	TotalHeight int `json:"totalHeight"`
	//TotalMoney  int `json:"totalMoney"`
}

type Transaction struct {
	UserID      int          `json:"userId"`
	AdminID     int          `json:"adminId"`
	TypeID      int          `json:"typeId"`
	Date        time.Time    `json:"date"`
	TotalMoney  int          `json:"totalMoney"`
	DepositID   int          `json:"depositID"`
	DataDeposit WasteDeposit `json:"dataDeposit"`
}

func ToDomainDeposit(req WasteDeposit) *transaction.DomainDeposit {
	return &transaction.DomainDeposit{
		ID:          req.ID,
		WasteId:     req.WasteId,
		TotalHeight: req.TotalHeight,
	}
}

func ToDomainTransaction(req Transaction) *transaction.DomainTransaction {
	return &transaction.DomainTransaction{
		TypeID:      req.TypeID,
		AdminID:     req.AdminID,
		Date:        req.Date,
		TotalMoney:  req.TotalMoney,
		DepositID:   req.DepositID,
		DepositData: *ToDomainDeposit(req.DataDeposit),
	}
}
