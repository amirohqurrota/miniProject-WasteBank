package users

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"wastebank-ca/app/middleware"
	"wastebank-ca/helper/encrypt"
)

type serviceUser struct {
	repository     Repository
	jwtAuth        *middleware.ConfigJWT
	contextTimeout time.Duration
}

func NewService(repoUser Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Service {
	return &serviceUser{
		repository:     repoUser,
		jwtAuth:        jwtauth,
		contextTimeout: timeout,
	}
}

func (servUser serviceUser) CreateToken(username, password string) (string, error) {
	// _, cancel := context.WithTimeout(ctx, servUser.contextTimeout)
	// defer cancel()

	if strings.TrimSpace(username) == "" && strings.TrimSpace(password) == "" {
		return "", errors.New("please fill username and password")
	}

	userDomain, err := servUser.GetData(0, "", "", username)
	if err != nil {
		return "", err
	}

	if !encrypt.ValidateHash(password, userDomain.Password) {
		return "", errors.New("invalid password")
	}

	token := servUser.jwtAuth.GenerateToken(userDomain.ID, "user")
	return token, nil
}

func (servUser serviceUser) Append(user *Domain) (*Domain, error) {
	//response,err:=servUser.Append()
	//fmt.Println(user.Username)
	result, err := servUser.repository.Insert(user)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servUser *serviceUser) Update(user *Domain) (*Domain, error) {
	//fmt.Println("id service", user.ID)
	result, err := servUser.repository.Update(user)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servUser *serviceUser) UpdateSaldo(id int, saldo int) (*Domain, error) {
	fmt.Println("update service")
	result, err := servUser.repository.UpdateSaldo(id, saldo)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servUser *serviceUser) GetData(id int, firstName string, lastName string, username string) (*Domain, error) {
	//fmt.Println("id service", id)
	result, err := servUser.repository.GetData(id, firstName, lastName, username)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}
