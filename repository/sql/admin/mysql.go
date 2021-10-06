package admin

import (
	"fmt"
	_adminDomain "wastebank-ca/bussines/admin"

	"gorm.io/gorm"
)

type repoAdmin struct {
	DBConn *gorm.DB
}

func NewRepoMySQL(db *gorm.DB) _adminDomain.Repository {
	return &repoAdmin{
		DBConn: db,
	}
}

func (repo *repoAdmin) Insert(admin *_adminDomain.Domain) (*_adminDomain.Domain, error) {
	recordUser := fromDomain(*admin)
	if err := repo.DBConn.Save(&recordUser).Error; err != nil {
		return &_adminDomain.Domain{}, err
	}
	result := toDomain(recordUser)
	return &result, nil
}

func (repo *repoAdmin) Update(admin *_adminDomain.Domain) (*_adminDomain.Domain, error) {
	recordAdmin := fromDomain(*admin)
	if err := repo.DBConn.Table("admins").Where("ID=?", admin.ID).Updates(recordAdmin).Error; err != nil {
		return &_adminDomain.Domain{}, err
	}
	var resultResponse Admin
	if err := repo.DBConn.Where("id=?", admin.ID).First(&resultResponse).Error; err != nil {
		return &_adminDomain.Domain{}, err
	}
	result := toDomain(resultResponse)
	return &result, nil
}
func (repo *repoAdmin) GetData(id int, name string) (*_adminDomain.Domain, error) {
	var recordAdmin Admin
	fmt.Println("id mysql ", id)
	if err := repo.DBConn.Where("first_name=? OR id=?", name, id).First(&recordAdmin).Error; err != nil {
		return &_adminDomain.Domain{}, err
	}

	result := toDomain(recordAdmin)
	return &result, nil
}
