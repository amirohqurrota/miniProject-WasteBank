package admin

import (
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

func (repo *repoAdmin) UpdateSaldo(id int, saldo int) (*_adminDomain.Domain, error) {
	var adminUpdateData Admin
	bonusPercentage := float32(0.1)
	if errFind := repo.DBConn.Where("id=?", id).First(&adminUpdateData).Error; errFind != nil {
		return &_adminDomain.Domain{}, errFind
	}
	adminUpdateData.TotalBonus += int(bonusPercentage * float32(saldo))
	if err := repo.DBConn.Table("users").Where("ID=?", id).Updates(adminUpdateData).Error; err != nil {
		return &_adminDomain.Domain{}, err
	}

	result := toDomain(adminUpdateData)
	return &result, nil
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

func (repo *repoAdmin) UpdateBonus(id int, total int) (*_adminDomain.Domain, error) {
	var adminUpdateData Admin
	percentageBonus := float32(0.1)
	if errFind := repo.DBConn.Where("id=?", id).First(&adminUpdateData).Error; errFind != nil {
		return &_adminDomain.Domain{}, errFind
	}
	adminUpdateData.TotalBonus += int(float32(total) * percentageBonus)
	if err := repo.DBConn.Table("users").Where("ID=?", id).Updates(adminUpdateData).Error; err != nil {
		return &_adminDomain.Domain{}, err
	}

	result := toDomain(adminUpdateData)
	return &result, nil
}

func (repo *repoAdmin) GetData(id int, firstName string, lastName string, username string) (*_adminDomain.Domain, error) {
	var recordAdmin Admin
	if err := repo.DBConn.Where("id=? OR first_name=? OR last_name=? OR username=?", id, firstName, lastName, username).First(&recordAdmin).Error; err != nil {
		return &_adminDomain.Domain{}, err
	}

	result := toDomain(recordAdmin)
	return &result, nil
}
