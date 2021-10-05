package request

import "wastebank-ca/bussines/users"

type UserInsert struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Telephone string `json:"telephone"`
	Address   string `json:"address"`
}

func ToDomain(request UserInsert) *users.Domain {
	return &users.Domain{
		ID:        request.ID,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Telephone: request.Telephone,
		Address:   request.Address,
	}
}

type UserGetById struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Telephone int    `json:"telephone"`
	Address   string `json:"address"`
}

type UpdateRequest struct {
	ID   int          `json:"id"`
	Data users.Domain `json:"data"`
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
