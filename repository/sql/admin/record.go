package admin

import (
	"wastebank-ca/bussines/admin"

	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	FirstName  string
	LastName   string
	Telephone  string
	Address    string
	TotalBonus int
}

func toDomain(rec Admin) admin.Domain {
	return admin.Domain{
		ID:         int(rec.ID),
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
		FirstName:  domain.FirstName,
		LastName:   domain.LastName,
		Telephone:  domain.Telephone,
		Address:    domain.Address,
		TotalBonus: domain.TotalBonus,
	}
}
