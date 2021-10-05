package response

import (
	"wastebank-ca/bussines/waste"
)

type WasteCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Waste struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	CategoryId    int    `json:"categoryId"`
	PurchasePrice int    `json:"purchasePrice"`
	TotalStock    int    `json:"totalStock"`
}

func FromDomain(domain waste.DomainWaste) Waste {
	return Waste{
		ID:            domain.ID,
		Name:          domain.Name,
		CategoryId:    domain.CategoryId,
		PurchasePrice: domain.PurchasePrice,
		TotalStock:    domain.TotalStock,
	}
}
