package waste

import (
	"time"
	"wastebank-ca/bussines/waste"

	"gorm.io/gorm"
)

type Waste struct {
	ID            int `gorm:"primaryKey"`
	Name          string
	CategoryId    int
	Category      WasteCategory `gorm:"foreignKey:CategoryId"`
	PurchasePrice int
	TotalStock    int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type WasteCategory struct {
	gorm.Model
	Name string
}

func toDomain(rec *Waste) waste.DomainWaste {
	return waste.DomainWaste{
		ID:            int(rec.ID),
		Name:          rec.Name,
		CategoryId:    int(rec.CategoryId),
		PurchasePrice: rec.PurchasePrice,
		TotalStock:    rec.TotalStock,
	}
}

func fromDomain(domain waste.DomainWaste) Waste {
	return Waste{
		ID:            domain.ID,
		Name:          domain.Name,
		CategoryId:    domain.CategoryId,
		PurchasePrice: domain.PurchasePrice,
		TotalStock:    domain.TotalStock,
	}
}
