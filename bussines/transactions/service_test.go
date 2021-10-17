package transactions_test

import (
	"os"
	"testing"
	_adminMocks "wastebank-ca/bussines/admin/mocks"
	_newsDomain "wastebank-ca/bussines/newsApi"
	_newsMocks "wastebank-ca/bussines/newsApi/mocks"
	"wastebank-ca/bussines/transactions"
	_transMocks "wastebank-ca/bussines/transactions/mocks"
	_userMocks "wastebank-ca/bussines/users/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	//adminServ _adminMocks.Service
	transService       transactions.Service
	transRepository    _transMocks.Repository
	adminService       _adminMocks.Service
	userService        _userMocks.Service
	newApiRepository   _newsMocks.Repository
	transDomain        transactions.DomainTransaction
	transDomainDeposit transactions.DomainDeposit
	newsDomain         _newsDomain.Domain
	transDomainType    transactions.DomainType
)

func TestMain(m *testing.M) {
	transService = transactions.NewService(&transRepository, &adminService, &userService, &newApiRepository)

	transDomainType = transactions.DomainType{
		ID:   1,
		Name: "organic",
	}

	transDomainDeposit = transactions.DomainDeposit{
		ID:          1,
		WasteId:     1,
		TotalHeight: 21,
	}
	transDomain = transactions.DomainTransaction{
		ID:         1,
		UserID:     1,
		AdminID:    1,
		TypeID:     1,
		TotalMoney: 3000,
		DepositID:  1,
	}

	newsDomain = _newsDomain.Domain{
		Source:      "cnn",
		Author:      "testname",
		Title:       "testTitle",
		Description: "desc",
		Content:     "ass",
		Url:         "wwww",
	}
	os.Exit(m.Run())

}

func TestAppend(t *testing.T) {
	t.Run("Append | Valid", func(t *testing.T) {

		userService.On("UpdateSaldo", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil, nil).Once()
		newApiRepository.On("GetNews").Return(newsDomain, nil).Once()
		transRepository.On("Insert", mock.Anything).Return(&transDomain, nil).Once()
		//newApiRepository.On()
		result, news, err := transService.Append(&transDomain)

		assert.Nil(t, err)
		assert.NotEmpty(t, result, news)
	})

	t.Run("Append | ValidWithTypeId2", func(t *testing.T) {

		userService.On("UpdateSaldo", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil, nil).Once()
		newApiRepository.On("GetNews").Return(newsDomain, nil).Once()
		transRepository.On("Insert", mock.Anything).Return(&transDomain, nil).Once()
		//newApiRepository.On()
		transDomainTest2 := transactions.DomainTransaction{
			ID:         1,
			UserID:     1,
			AdminID:    1,
			TypeID:     2,
			TotalMoney: 3000,
			DepositID:  1,
		}
		result, news, err := transService.Append(&transDomainTest2)

		assert.Nil(t, err)
		assert.NotEmpty(t, result, news)
	})

	t.Run("Append | InvalidUpdateSaldo", func(t *testing.T) {

		userService.On("UpdateSaldo", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(nil, assert.AnError).Once()
		//newApiRepository.On("GetNews").Return(newsDomain, nil).Once()
		//transRepository.On("Insert", mock.Anything).Return(&transDomain, nil).Once()
		//newApiRepository.On()
		result, news, err := transService.Append(&transDomain)

		assert.NotNil(t, err)
		assert.Empty(t, result, news)
	})

}

func TestAddNewType(t *testing.T) {
	t.Run("AddNewType | Valid", func(t *testing.T) {
		transRepository.On("AddNewType", mock.Anything).Return(&transDomainType, nil).Once()
		result, err := transService.AddNewType(&transDomainType)

		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("AddNewType | Invalid", func(t *testing.T) {
		transRepository.On("AddNewType", mock.Anything).Return(&transDomainType, assert.AnError).Once()
		result, err := transService.AddNewType(&transDomainType)

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})

}
