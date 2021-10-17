package request

import (
	"time"
	"wastebank-ca/bussines/admin"
)

type Admin struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Telephone  string `json:"telephone"`
	Address    string `json:"address"`
	TotalBonus int    `json:"totalBonus"`
}

func ToDomain(request Admin) *admin.Domain {
	return &admin.Domain{
		ID:         request.ID,
		Username:   request.Username,
		Password:   request.Password,
		FirstName:  request.FirstName,
		LastName:   request.LastName,
		Telephone:  request.Telephone,
		Address:    request.Address,
		TotalBonus: request.TotalBonus,
	}
}

type DataUpdate struct {
	FirstName  string
	LastName   string
	Telephone  string
	Address    string
	TotalBonus int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type UpdateRequest struct {
	ID   int        `json:"id"`
	Data DataUpdate `json:"data"`
}

func UpdateToDomain(request UpdateRequest) *admin.Domain {
	return &admin.Domain{
		ID:         request.ID,
		FirstName:  request.Data.FirstName,
		LastName:   request.Data.LastName,
		Telephone:  request.Data.Telephone,
		Address:    request.Data.Address,
		TotalBonus: request.Data.TotalBonus,
	}
}
