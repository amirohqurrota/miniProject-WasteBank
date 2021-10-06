package response

import "wastebank-ca/bussines/admin"

type Admin struct {
	ID         int    `json:"id"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Telephone  string `json:"telephone"`
	Address    string `json:"address"`
	TotalBonus int    `json:"totalBonus"`
}

func FromDomain(domain admin.Domain) Admin {
	return Admin{
		ID:         domain.ID,
		FirstName:  domain.FirstName,
		LastName:   domain.LastName,
		Telephone:  domain.Telephone,
		Address:    domain.Address,
		TotalBonus: domain.TotalBonus,
	}
}
