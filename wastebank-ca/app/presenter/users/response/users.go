package response

import "wastebank-ca/bussines/users"

type Users struct {
	ID         int    `json:"id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Telephone  string `json:"telephone"`
	Address    string `json:"address"`
	TotalWaste int    `json:"totalWaste"`
	TotalSaldo int    `json:"totalSaldo"`
}

func FromDomain(domain users.Domain) Users {
	return Users{
		ID:         domain.ID,
		FirstName:  domain.FirstName,
		LastName:   domain.LastName,
		Telephone:  domain.Telephone,
		Address:    domain.Address,
		TotalWaste: domain.TotalWaste,
		TotalSaldo: domain.TotalSaldo,
	}
}
