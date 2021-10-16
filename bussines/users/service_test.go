package users_test

import (
	"testing"
	"time"
	"wastebank-ca/bussines/users"
	_userMocks "wastebank-ca/bussines/users/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	//adminServ _adminMocks.Service
	userService    users.Service
	userRepository _userMocks.Repository
	userDomain     users.Domain
	//userDomainUpdate    users.Domain
)

func TestMain(m *testing.M) {
	userService = users.NewService(&userRepository, nil, time.Hour*1)

	userDomain = users.Domain{
		ID:         1,
		Username:   "nama",
		Password:   "pass",
		FirstName:  "firstName",
		LastName:   "lastName",
		TotalWaste: 20,
		TotalSaldo: 10000,
	}

}

func TestAppend(t *testing.T) {
	t.Run("Append | Valid", func(t *testing.T) {
		//adminService.On("Append", mock.Anything, mock.AnythingOfType("*Domain")).Return(adminDomain, nil).Once()
		userRepository.On("Insert", mock.Anything, mock.AnythingOfType("*Domain")).Return(userDomain, nil).Once()

		result, err := userService.Append(&userDomain)

		assert.Nil(t, err)
		assert.Contains(t, "nama", result.Username)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		userRepository.On("Update", mock.AnythingOfType("*Domain")).Return(userDomain, nil).Once()
		result, err := userRepository.Update(&userDomain)
		assert.Nil(t, err)
		assert.Contains(t, "nama", result.Username)

	})
}
