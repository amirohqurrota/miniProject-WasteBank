package response

import (
	"time"
	newsapi "wastebank-ca/bussines/newsApi"
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

//response for transaction get data in general
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

//response for insert new transaction
type News struct {
	Source      string `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"desc"`
	Content     string `json:"content"`
	Url         string `json:"url"`
}

func FromDomainNews(domain newsapi.Domain) News {
	return News{
		Source:      domain.Source,
		Author:      domain.Author,
		Title:       domain.Title,
		Description: domain.Description,
		Content:     domain.Content,
		Url:         domain.Url,
	}
}

type AddNewTransaction struct {
	ID          int       `json:"id"`
	UserID      int       `json:"userId"`
	AdminID     int       `json:"adminId"`
	TypeID      int       `json:"typeId"`
	Date        time.Time `json:"date"`
	TotalMoney  int       `json:"totalMoney"`
	DepositID   int       `json:"depositID"`
	DataDeposit []Deposit `json:"dataDeposit"`
	News        News      `json:"news"`
}

func FromDomainNewTrans(domain transactions.DomainTransaction, domainNews newsapi.Domain) AddNewTransaction {
	if domain.TypeID == 1 {
		//fmt.Println(domain.DepositData[0])
		return AddNewTransaction{
			TypeID:      domain.TypeID,
			ID:          domain.ID,
			UserID:      domain.UserID,
			AdminID:     domain.AdminID,
			TotalMoney:  domain.TotalMoney,
			Date:        domain.Date,
			DataDeposit: fromDomainAllDeposit(domain.DepositData),
			News:        FromDomainNews(domainNews),
		}
	}
	return AddNewTransaction{
		TypeID:     domain.TypeID,
		ID:         domain.ID,
		UserID:     domain.UserID,
		AdminID:    domain.AdminID,
		TotalMoney: domain.TotalMoney,
		Date:       domain.Date,
		News:       FromDomainNews(domainNews),
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

// var Quotes = []string{
// 	"The ultimate test of man’s conscience may be his willingness to sacrificesomething today for future generations whose words of thanks will not be heard.",
// 	"The Earth is a fine place and worth fighting for",
// }

// var Quotes = []string{
// 	"Thank you for taking care and saving the planet with small action in sorting waste. because the earth is a fine place, so it worth to fighting for",
// }
