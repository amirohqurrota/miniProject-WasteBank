package users

import (
	"wastebank-ca/bussines/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName  string
	LastName   string
	Telephone  string
	Address    string
	TotalWaste int
	TotalSaldo int
}

func toDomain(rec User) users.Domain {
	return users.Domain{
		ID:         int(rec.ID),
		FirstName:  rec.FirstName,
		LastName:   rec.LastName,
		Telephone:  rec.Telephone,
		Address:    rec.Address,
		TotalWaste: rec.TotalWaste,
		TotalSaldo: rec.TotalSaldo,
		CreatedAt:  rec.CreatedAt,
		UpdatedAt:  rec.UpdatedAt,
	}
}

func fromDomain(domain users.Domain) User {
	return User{
		FirstName:  domain.FirstName,
		LastName:   domain.LastName,
		Telephone:  domain.Telephone,
		Address:    domain.Address,
		TotalWaste: domain.TotalWaste,
		TotalSaldo: domain.TotalSaldo,
	}
}
