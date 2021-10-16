package admin

import (
	"errors"
	"fmt"
	"time"
	"wastebank-ca/app/middleware"
	"wastebank-ca/helper/encrypt"
)

type serviceAdmin struct {
	repository Repository
	jwtAuth    *middleware.ConfigJWT
	// contextTimeout time.Duration
}

func NewService(repoAdmin Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Service {
	return &serviceAdmin{
		repository: repoAdmin,
		jwtAuth:    jwtauth,
		// contextTimeout: timeout,
	}
}

func (servAdmin serviceAdmin) CreateToken(username, password string) (string, error) {
	// _, cancel := context.WithTimeout(ctx, servAdmin.contextTimeout)
	// defer cancel()

	// if strings.TrimSpace(username) == "" && strings.TrimSpace(password) == "" {
	// 	return "", errors.New("please fill username and password")
	// }

	adminDomain, err := servAdmin.GetData(0, "", "", username)
	if err != nil {
		return "", err
	}

	if !encrypt.ValidateHash(password, adminDomain.Password) {
		return "", errors.New("invalid password")
	}

	token := servAdmin.jwtAuth.GenerateToken(adminDomain.ID, "admin")
	return token, nil
}

func (servAdmin serviceAdmin) Append(admin *Domain) (*Domain, error) {
	passwordHashed := encrypt.Hash(admin.Password)
	admin.Password = passwordHashed
	result, err := servAdmin.repository.Insert(admin)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servAdmin *serviceAdmin) Update(admin *Domain) (*Domain, error) {
	//fmt.Println("id service", admin.ID)
	result, err := servAdmin.repository.Update(admin)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servUser *serviceAdmin) UpdateSaldo(id int, saldo int) (*Domain, error) {
	result, err := servUser.repository.UpdateSaldo(id, saldo)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servAdmin *serviceAdmin) GetData(id int, firstName string, lastName string, username string) (*Domain, error) {
	fmt.Println("id service", id)
	result, err := servAdmin.repository.GetData(id, firstName, lastName, username)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}
