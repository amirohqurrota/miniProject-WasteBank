package response

import (
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

func FromDomain(domain admin.Domain) Admin {
	return Admin{
		ID:         domain.ID,
		Username:   domain.Username,
		Password:   domain.Password,
		FirstName:  domain.FirstName,
		LastName:   domain.LastName,
		Telephone:  domain.Telephone,
		Address:    domain.Address,
		TotalBonus: domain.TotalBonus,
	}
}

type TokenResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func FromDomainToken(token string, domain admin.Domain) TokenResponse {
	return TokenResponse{
		Username: domain.Username,
		Token:    token,
	}
}
