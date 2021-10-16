package admin_test

import (
	"os"
	"testing"
	"time"
	"wastebank-ca/app/middleware"
	"wastebank-ca/bussines/admin"
	_adminMocks "wastebank-ca/bussines/admin/mocks"
	"wastebank-ca/helper/encrypt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	//adminServ _adminMocks.Service
	adminService    admin.Service
	adminRepository _adminMocks.Repository
	adminDomainTest admin.Domain
	passwordHashed  string
)

func TestMain(m *testing.M) {
	adminService = admin.NewService(&adminRepository, &middleware.ConfigJWT{}, time.Hour*1)
	passwordHashed = encrypt.Hash("password")

	adminDomainTest = admin.Domain{
		ID:         1,
		Username:   "nama",
		Password:   passwordHashed,
		FirstName:  "firstName",
		LastName:   "lastName",
		TotalBonus: 20000,
	}
	os.Exit(m.Run())
}

func TestAppend(t *testing.T) {
	t.Run("Append | Valid", func(t *testing.T) {
		//adminService.On("Append", mock.Anything, mock.AnythingOfType("*Domain")).Return(adminDomain, nil).Once()
		adminRepository.On("Insert", mock.Anything).Return(&adminDomainTest, nil).Once()

		result, err := adminService.Append(&adminDomainTest)

		assert.Nil(t, err)
		assert.Contains(t, "nama", result.Username)
	})

}

func TestCreateToken(t *testing.T) {
	// t.Run("CreateToken | Valid", func(t *testing.T) {
	// 	adminRepository.On("GetData", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&adminDomainTest, nil).Once()

	// 	result, err := adminService.CreateToken("nama", "password")
	// 	assert.Nil(t, err)
	// 	assert.NotEmpty(t, result)
	// 	//assert.Contains(t, "nama", result.Username)
	// })

	t.Run("CreateToken | Invalid Password", func(t *testing.T) {
		adminRepository.On("GetData", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&adminDomainTest, assert.AnError).Once()
		adminRepository.On("GenerateToken", mock.AnythingOfType("int"), "admin").Return()
		result, err := adminService.CreateToken("nama", "passwordSalah")
		assert.NotNil(t, err)
		assert.Empty(t, result)
		//assert.Contains(t, "nama", result.Username)
	})

	t.Run("CreateToken | Invalid username", func(t *testing.T) {
		adminRepository.On("GetData", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&adminDomainTest, nil).Once()
		adminRepository.On("GenerateToken", mock.AnythingOfType("int"), "admin").Return()
		result, err := adminService.CreateToken("bukanNama", "passwordSalah")
		assert.NotNil(t, err)
		assert.Empty(t, result)
		//assert.Contains(t, "nama", result.Username)
	})
}

func TestGetData(t *testing.T) {
	t.Run("GetData | Valid", func(t *testing.T) {
		adminRepository.On("GetData", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(&adminDomainTest, nil).Once()
		result, err := adminService.GetData(1, "", "", "")
		assert.Nil(t, err)
		assert.Contains(t, "nama", result.Username)
	})

	t.Run("GetData | Invalid", func(t *testing.T) {
		adminRepository.On("GetData", mock.AnythingOfType("int"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(nil, assert.AnError).Once()
		result, err := adminService.GetData(3, "wrongFirstName", "", "")
		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		adminRepository.On("Update", mock.Anything).Return(&adminDomainTest, nil).Once()
		updateData := admin.Domain{
			ID:        1,
			FirstName: "FirstNamaUpdate",
		}
		result, err := adminService.Update(&updateData)
		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Update | Invalid", func(t *testing.T) {
		adminRepository.On("Update", mock.Anything).Return(&adminDomainTest, assert.AnError).Once()
		updateData := admin.Domain{
			ID:        8,
			FirstName: "FirstNamaUpdate",
		}
		result, err := adminService.Update(&updateData)
		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestUpdateSaldo(t *testing.T) {
	t.Run("UpdateSaldo | Valid", func(t *testing.T) {
		adminRepository.On("UpdateSaldo", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(&adminDomainTest, nil).Once()

		result, err := adminService.UpdateSaldo(1, 25000)
		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("UpdateSaldo | Invalid", func(t *testing.T) {
		adminRepository.On("UpdateSaldo", mock.AnythingOfType("int"), mock.AnythingOfType("int")).Return(&adminDomainTest, assert.AnError).Once()

		result, err := adminService.UpdateSaldo(8, 25000)
		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}
