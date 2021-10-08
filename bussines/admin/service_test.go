package admin_test

import (
	"testing"
	"time"
	"wastebank-ca/bussines/admin"
	_adminMocks "wastebank-ca/bussines/admin/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	//adminServ _adminMocks.Service
	adminService    admin.Service
	adminRepository _adminMocks.Repository
	adminDomain     admin.Domain
)

func TestMain(m *testing.M) {
	adminService = admin.NewService(&adminRepository, nil, time.Hour*1)

	adminDomain = admin.Domain{
		ID:         1,
		Username:   "nama",
		Password:   "pass",
		FirstName:  "firstName",
		LastName:   "lastName",
		TotalBonus: 20000,
	}

}

func TestAppend(t *testing.T) {
	t.Run("Append | Valid", func(t *testing.T) {
		//adminService.On("Append", mock.Anything, mock.AnythingOfType("*Domain")).Return(adminDomain, nil).Once()
		adminRepository.On("Insert", mock.Anything, mock.AnythingOfType("*Domain")).Return(adminDomain, nil).Once()

		result, err := adminService.Append(&adminDomain)

		assert.Nil(t, err)
		assert.Contains(t, "nama", result.Username)
	})
}

func TestCreateToken(t *testing.T) {
	t.Run("CreateToken | Valid", func(t *testing.T) {
		adminRepository.On("GetData", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(adminDomain, nil).Once()
		result, err := adminService.GetData(1, "nama", "depan", "username")
		adminRepository.On("GenerateToken", 1, "admin")
		//token:=adminService.jwtAuth.GenerateToken()
		assert.Nil(t, err)
		assert.Contains(t, "nama", result.Username)
	})
}

func TestGetData(t *testing.T) {
	t.Run("GetData | Valid", func(t *testing.T) {
		adminRepository.On("GetData", mock.Anything, mock.Anything, mock.Anything).Return(adminDomain, nil).Once()
		result, err := adminRepository.GetData(1, "", "", "")
		assert.Nil(t, err)
		assert.Contains(t, "nama", result.Username)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		adminRepository.On("Update", mock.AnythingOfType("*Domain")).Return(adminDomain, nil).Once()
		result, err := adminRepository.Update(&adminDomain)
		assert.Nil(t, err)
		assert.Contains(t, "nama", result.Username)
	})
}
