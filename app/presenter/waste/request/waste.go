package request

import (
	"wastebank-ca/bussines/waste"
)

type WasteInsert struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	CategoryId    int    `json:"categoryId"`
	PurchasePrice int    `json:"purchasePrice"`
	TotalStock    int    `json:"totalStock"`
}

func ToDomainInsert(request WasteInsert) *waste.DomainWaste {
	return &waste.DomainWaste{
		ID:            request.ID,
		Name:          request.Name,
		CategoryId:    request.CategoryId,
		PurchasePrice: request.PurchasePrice,
		TotalStock:    request.TotalStock,
	}
}

type UpdateRequest struct {
	ID   int               `json:"id"`
	Data waste.DomainWaste `json:"data"`
}

func UpdateToDomain(request UpdateRequest) *waste.DomainWaste {
	return &waste.DomainWaste{
		ID:            request.ID,
		Name:          request.Data.Name,
		CategoryId:    request.Data.CategoryId,
		TotalStock:    request.Data.TotalStock,
		PurchasePrice: request.Data.PurchasePrice,
	}
}
