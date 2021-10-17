package waste_test

import (
	"os"
	"testing"
	"wastebank-ca/bussines/waste"
	_wasteMocks "wastebank-ca/bussines/waste/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	wasteService        waste.Service
	wasteRepository     _wasteMocks.Repository
	wasteDomain         waste.DomainWaste
	wasteDomainCategory waste.DomainCategory
	listOfWaste         []waste.DomainWaste
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

	listOfWaste = []waste.DomainWaste{
		{
			ID:            1,
			Name:          "paper",
			CategoryId:    1,
			TotalStock:    23,
			PurchasePrice: 1000,
		},
		{
			ID:            2,
			Name:          "paper2",
			CategoryId:    1,
			TotalStock:    23,
			PurchasePrice: 1000,
		},
	}
	os.Exit(m.Run())

}

func TestAppend(t *testing.T) {
	t.Run("Append | Valid", func(t *testing.T) {
		wasteRepository.On("Insert", mock.Anything, mock.Anything).Return(&wasteDomain, nil).Once()

		result, err := wasteService.Append(&wasteDomain)

		assert.Nil(t, err)
		assert.Contains(t, result.Name, "paper")
	})

	t.Run("Append | Invalid", func(t *testing.T) {
		wasteRepository.On("Insert", mock.Anything, mock.Anything).Return(&wasteDomain, assert.AnError).Once()

		_, err := wasteService.Append(&wasteDomain)

		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		wasteRepository.On("Update", mock.Anything).Return(&wasteDomain, nil).Once()
		result, err := wasteService.Update(&wasteDomain)
		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Update | Invalid", func(t *testing.T) {
		wasteRepository.On("Update", mock.Anything).Return(&wasteDomain, assert.AnError).Once()
		result, err := wasteService.Update(&wasteDomain)
		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestFindAll(t *testing.T) {
	t.Run("FindAll | Valid", func(t *testing.T) {
		wasteRepository.On("FindAll").Return(&listOfWaste, nil).Once()
		result, err := wasteService.FindAll()
		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("FindAll | Invalid", func(t *testing.T) {
		wasteRepository.On("FindAll").Return(&listOfWaste, assert.AnError).Once()
		result, err := wasteService.FindAll()
		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestGetData(t *testing.T) {
	t.Run("GetData | Valid", func(t *testing.T) {
		wasteRepository.On("GetData", mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(&wasteDomain, nil).Once()
		result, err := wasteService.GetData(1, "paper")
		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("GetData | Invalid", func(t *testing.T) {
		wasteRepository.On("GetData", mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(&wasteDomain, assert.AnError).Once()
		result, err := wasteService.GetData(1, "wronginput")
		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}
