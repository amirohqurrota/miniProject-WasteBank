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
	//fmt.Println("id mysql ", user.ID)
	if err := repo.DBConn.Table("users").Where("ID=?", user.ID).Updates(recordUser).Error; err != nil {
		return &users.Domain{}, err
	}
	var resultResponse User
	if err := repo.DBConn.Where("id=?", user.ID).First(&resultResponse).Error; err != nil {
		return &users.Domain{}, err
	}
	result := toDomain(resultResponse)
	return &result, nil
}

func (repo *repoUser) UpdateSaldo(id int, saldo int) (*users.Domain, error) {
	//recordUser := fromDomain(*user)
	fmt.Println("id mysql ", id)
	var userUpdateData User
	if errFind := repo.DBConn.Where("id=?", id).First(&userUpdateData).Error; errFind != nil {
		fmt.Println("error find mysql ")
		return &users.Domain{}, errFind
	}
	fmt.Println("error find mysql passed ")
	userUpdateData.TotalSaldo += saldo
	//fmt.Println("id mysql ", user.ID)
	if err := repo.DBConn.Table("users").Where("ID=?", id).Updates(userUpdateData).Error; err != nil {
		fmt.Println("error update mysql ")
		return &users.Domain{}, err
	}

	result := toDomain(userUpdateData)
	return &result, nil
}

func (repo *repoUser) GetData(id int, name string) (*users.Domain, error) {
	var recordUser User
	//fmt.Println("id mysql ", id)
	if err := repo.DBConn.Where("first_name=? OR id=?", name, id).First(&recordUser).Error; err != nil {
		return &users.Domain{}, err
	}

	result := toDomain(recordUser)
	return &result, nil
}
