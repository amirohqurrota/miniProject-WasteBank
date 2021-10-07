package request

import (
	"time"
	"wastebank-ca/bussines/users"
)

type UserInsert struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Telephone  string `json:"telephone"`
	Address    string `json:"address"`
	TotalWaste int    `json:"totalWaste"`
	TotalSaldo int    `json:"totalSaldo"`
}

func ToDomain(request UserInsert) *users.Domain {
	return &users.Domain{
		ID:         request.ID,
		Username:   request.Username,
		Password:   request.Password,
		FirstName:  request.FirstName,
		LastName:   request.LastName,
		Telephone:  request.Telephone,
		Address:    request.Address,
		TotalWaste: request.TotalWaste,
		TotalSaldo: request.TotalSaldo,
	}
}

type UserGetById struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Telephone int    `json:"telephone"`
	Address   string `json:"address"`
}

type DataUpdate struct {
	FirstName  string
	LastName   string
	Telephone  string
	Address    string
	TotalWaste int
	TotalSaldo int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type UpdateRequest struct {
	ID   int        `json:"id"`
	Data DataUpdate `json:"data"`
}

func UpdateToDomain(request UpdateRequest) *users.Domain {
	return &users.Domain{
		ID:         request.ID,
		FirstName:  request.Data.FirstName,
		LastName:   request.Data.LastName,
		Telephone:  request.Data.Telephone,
		Address:    request.Data.Address,
		TotalWaste: request.Data.TotalWaste,
		TotalSaldo: request.Data.TotalSaldo,
	}
}
