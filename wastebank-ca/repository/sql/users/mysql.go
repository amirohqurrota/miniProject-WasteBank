package users

import (
	"fmt"
	"wastebank-ca/bussines/users"

	"gorm.io/gorm"
)

type repoUser struct {
	DBConn *gorm.DB
}

func NewRepoMySQL(db *gorm.DB) users.Repository {
	return &repoUser{
		DBConn: db,
	}
}

func (repo *repoUser) Insert(user *users.Domain) (*users.Domain, error) {
	recordUser := fromDomain(*user)
	if err := repo.DBConn.Save(&recordUser).Error; err != nil {
		return &users.Domain{}, err
	}
	result := toDomain(recordUser)
	return &result, nil
}

func (repo *repoUser) Update(user *users.Domain) (*users.Domain, error) {
	recordUser := fromDomain(*user)
	fmt.Println("id mysql ", user.ID)
	if err := repo.DBConn.Table("users").Where("ID=?", user.ID).Updates(user).Error; err != nil {
		return &users.Domain{}, err
	}
	result := toDomain(recordUser)
	return &result, nil
}
func (repo *repoUser) FindByID(id int) (*users.Domain, error) {
	var recordUser User
	if err := repo.DBConn.First(&recordUser, id).Error; err != nil {
		return &users.Domain{}, err
	}
	result := toDomain(recordUser)
	return &result, nil
}
