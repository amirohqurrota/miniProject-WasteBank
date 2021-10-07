package admin

import (
	"wastebank-ca/bussines/admin"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	username   string
	password   string
	FirstName  string
	LastName   string
	Telephone  string
	Address    string
	TotalBonus int
}

func toDomain(rec Admin) admin.Domain {
	return admin.Domain{
		ID:         int(rec.ID),
		Username:   rec.username,
		Password:   rec.password,
		FirstName:  rec.FirstName,
		LastName:   rec.LastName,
		Telephone:  rec.Telephone,
		Address:    rec.Address,
		TotalBonus: rec.TotalBonus,
		CreatedAt:  rec.CreatedAt,
		UpdatedAt:  rec.UpdatedAt,
	}
}

func fromDomain(domain admin.Domain) Admin {
	return Admin{
		username:   domain.Username,
		password:   domain.Password,
		FirstName:  domain.FirstName,
		LastName:   domain.LastName,
		Telephone:  domain.Telephone,
		Address:    domain.Address,
		TotalBonus: domain.TotalBonus,
	}
}
