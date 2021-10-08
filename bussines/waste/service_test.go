package waste_test

import (
	"testing"
	"wastebank-ca/bussines/waste"
	_wasteMocks "wastebank-ca/bussines/waste/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	//adminServ _adminMocks.Service
	wasteService        waste.Service
	wasteRepository     _wasteMocks.Repository
	wasteDomain         waste.DomainWaste
	wasteDomainCategory waste.DomainCategory
	//userDomainUpdate    users.Domain
)

func TestMain(m *testing.M) {
	wasteService = waste.NewService(&wasteRepository)

	wasteDomain = waste.DomainWaste{
		ID:            1,
		Name:          "paper",
		CategoryId:    1,
		TotalStock:    23,
		PurchasePrice: 1000,
	}
	wasteDomainCategory = waste.DomainCategory{
		ID:   1,
		Name: "organic",
	}

}

func TestAppend(t *testing.T) {
	t.Run("Append | Valid", func(t *testing.T) {
		//adminService.On("Append", mock.Anything, mock.AnythingOfType("*Domain")).Return(adminDomain, nil).Once()
		wasteRepository.On("Insert", mock.Anything, mock.AnythingOfType("*DomainWaste")).Return(wasteDomain, nil).Once()

		result, err := wasteService.Append(&wasteDomain)

		assert.Nil(t, err)
		assert.Contains(t, "paper", result.Name)
	})
}

// func TestUpdate(t *testing.T) {
// 	t.Run("Update | Valid", func(t *testing.T) {
// 		userRepository.On("Update", mock.AnythingOfType("*Domain")).Return(userDomain, nil).Once()
// 		result, err := userRepository.Update(&userDomain)
// 		assert.Nil(t, err)
// 		assert.Contains(t, "nama", result.Username)
// 	})
// }
