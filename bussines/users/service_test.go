package users_test

import (
	"os"
	"testing"
	"time"
	"wastebank-ca/bussines/users"
	_userMocks "wastebank-ca/bussines/users/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	//userServ _userMocks.Service
	userService    users.Service
	userRepository _userMocks.Repository
	userDomainTest users.Domain
	//userDomainUpdate    users.Domain
)

func TestMain(m *testing.M) {
	userService = users.NewService(&userRepository, nil, time.Hour*1)

	userDomainTest = users.Domain{
		ID:         1,
		Username:   "nama",
		Password:   "pass",
		FirstName:  "firstName",
		LastName:   "lastName",
		TotalWaste: 20,
		TotalSaldo: 10000,
	}
	os.Exit(m.Run())

}

func TestAppend(t *testing.T) {
	t.Run("Append | Valid", func(t *testing.T) {
		//userService.On("Append", mock.Anything, mock.AnythingOfType("*Domain")).Return(userDomain, nil).Once()
		userRepository.On("Insert", mock.Anything).Return(&userDomainTest, nil).Once()

		result, err := userService.Append(&userDomainTest)

		assert.Nil(t, err)
		assert.Contains(t, "nama", result.Username)
	})

}

func TestCreateToken(t *testing.T) {
	// t.Run("CreateToken | Valid", func(t *testing.T) {
	// 	userRepository.On("GetData", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&userDomainTest, nil).Once()

	// 	result, err := userService.CreateToken("nama", "password")
	// 	assert.Nil(t, err)
	// 	assert.NotEmpty(t, result)
	// 	//assert.Contains(t, "nama", result.Username)
	// })

	t.Run("CreateToken | Invalid Password", func(t *testing.T) {
		userRepository.On("GetData", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&userDomainTest, assert.AnError).Once()
		userRepository.On("GenerateToken", mock.AnythingOfType("int"), "user").Return()
		result, err := userService.CreateToken("nama", "passwordSalah")
		assert.NotNil(t, err)
		assert.Empty(t, result)
		//assert.Contains(t, "nama", result.Username)
	})

	t.Run("CreateToken | Invalid username", func(t *testing.T) {
		userRepository.On("GetData", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&userDomainTest, nil).Once()
		userRepository.On("GenerateToken", mock.AnythingOfType("int"), "user").Return()
		result, err := userService.CreateToken("bukanNama", "passwordSalah")
		assert.NotNil(t, err)
		assert.Empty(t, result)
		//assert.Contains(t, "nama", result.Username)
	})
}

func TestGetData(t *testing.T) {
	t.Run("GetData | Valid", func(t *testing.T) {
		userRepository.On("GetData", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&userDomainTest, nil).Once()
		result, err := userService.GetData(1, "", "", "")
		assert.Nil(t, err)
		assert.Contains(t, "nama", result.Username)
	})

	t.Run("GetData | Invalid", func(t *testing.T) {
		userRepository.On("GetData", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, assert.AnError).Once()
		result, err := userService.GetData(3, "wrongFirstName", "", "")
		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		userRepository.On("Update", mock.Anything).Return(&userDomainTest, nil).Once()
		updateData := users.Domain{
			ID:        1,
			FirstName: "FirstNamaUpdate",
		}
		result, err := userService.Update(&updateData)
		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Update | Invalid", func(t *testing.T) {
		userRepository.On("Update", mock.Anything).Return(&userDomainTest, assert.AnError).Once()
		updateData := users.Domain{
			ID:        8,
			FirstName: "FirstNamaUpdate",
		}
		result, err := userService.Update(&updateData)
		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestUpdateSaldo(t *testing.T) {
	t.Run("UpdateSaldo | Valid", func(t *testing.T) {
		userRepository.On("UpdateSaldo", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(&userDomainTest, nil).Once()

		result, err := userService.UpdateSaldo(1, 25000)
		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("UpdateSaldo | Invalid", func(t *testing.T) {
		userRepository.On("UpdateSaldo", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(&userDomainTest, assert.AnError).Once()

		result, err := userService.UpdateSaldo(8, 25000)
		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}
