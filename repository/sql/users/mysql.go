package users

import (
	"fmt"
	"wastebank-ca/bussines/users"

	"golang.org/x/crypto/bcrypt"
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

func Hashpassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }

func (repo *repoUser) Insert(user *users.Domain) (*users.Domain, error) {
	user.Password, _ = Hashpassword(user.Password)
	recordUser := fromDomain(*user)
	fmt.Println(recordUser.Username)
	if err := repo.DBConn.Save(&recordUser).Error; err != nil {
		return &users.Domain{}, err
	}
	result := toDomain(recordUser)
	return &result, nil
}

func (repo *repoUser) Update(user *users.Domain) (*users.Domain, error) {
	recordUser := fromDomain(*user)
	//fmt.Println("id mysql ", user.ID)
	if err := repo.DBConn.Table("users").Where("id=?", user.ID).Updates(recordUser).Error; err != nil {
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
	var userUpdateData User
	if errFind := repo.DBConn.Where("id=?", id).First(&userUpdateData).Error; errFind != nil {
		fmt.Println("error find mysql ")
		return &users.Domain{}, errFind
	}
	//fmt.Println("error find mysql passed ")
	userUpdateData.TotalSaldo += saldo
	if err := repo.DBConn.Table("users").Where("ID=?", id).Updates(userUpdateData).Error; err != nil {
		fmt.Println("error update mysql ")
		return &users.Domain{}, err
	}

	result := toDomain(userUpdateData)
	return &result, nil
}

func (repo *repoUser) GetData(id int, firstName string, lastName string, username string) (*users.Domain, error) {
	var recordUser User
	if err := repo.DBConn.Where("id=? OR first_name=? OR last_name=? OR username=?", id, firstName, lastName, username).First(&recordUser).Error; err != nil {
		return &users.Domain{}, err
	}
	result := toDomain(recordUser)
	return &result, nil
}
