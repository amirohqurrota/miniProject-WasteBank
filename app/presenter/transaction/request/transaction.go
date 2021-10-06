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
}

type Transaction struct {
	UserID      int            `json:"userId"`
	AdminID     int            `json:"adminId"`
	TypeID      int            `json:"typeId"`
	Date        time.Time      `json:"date"`
	TotalMoney  int            `json:"totalMoney"`
	DepositID   int            `json:"depositID"`
	DataDeposit []WasteDeposit `json:"dataDeposit"`
}

type TypeTransaction struct {
	//ID   int    `json:"id"`
	Name string `json:"name"`
}

func ToDomainTypeTrans(req TypeTransaction) *transaction.DomainType {
	return &transaction.DomainType{
		Name: req.Name,
	}
}

func ToDomainDeposit(req WasteDeposit) *transaction.DomainDeposit {
	return &transaction.DomainDeposit{
		ID:          req.ID,
		WasteId:     req.WasteId,
		TotalHeight: req.TotalHeight,
	}
}

func ToDomainDepositAll(req []WasteDeposit) *[]transaction.DomainDeposit {
	var arrayDeposit []transaction.DomainDeposit
	for _, n := range req {
		arrayDeposit = append(arrayDeposit, *ToDomainDeposit(n))
	}
	return &arrayDeposit
}

func ToDomainTransaction(req Transaction) *transaction.DomainTransaction {
	return &transaction.DomainTransaction{
		UserID:      req.UserID,
		AdminID:     req.AdminID,
		TypeID:      req.TypeID,
		Date:        req.Date,
		TotalMoney:  req.TotalMoney,
		DepositID:   req.DepositID,
		DepositData: *ToDomainDepositAll(req.DataDeposit),
	}
}
