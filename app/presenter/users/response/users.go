package response

import "wastebank-ca/bussines/users"

type Users struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Telephone  string `json:"telephone"`
	Address    string `json:"address"`
	TotalWaste int    `json:"totalWaste"`
	TotalSaldo int    `json:"totalSaldo"`
	Token      string `json:"token"`
}

func FromDomain(domain users.Domain) Users {
	return Users{
		ID:         domain.ID,
		Username:   domain.Username,
		FirstName:  domain.FirstName,
		LastName:   domain.LastName,
		Telephone:  domain.Telephone,
		Address:    domain.Address,
		TotalWaste: domain.TotalWaste,
		TotalSaldo: domain.TotalSaldo,
	}
}

type TokenResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func FromDomainToken(token string, domain users.Domain) TokenResponse {
	return TokenResponse{
		Username: domain.Username,
		Token:    token,
	}
}
