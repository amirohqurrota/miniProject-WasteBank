package users

import "fmt"

type serviceUser struct {
	repository Repository
}

func NewService(repoUser Repository) Service {
	return &serviceUser{
		repository: repoUser,
	}
}

func (servUser serviceUser) Append(user *Domain) (*Domain, error) {
	//response,err:=servUser.Append()
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

func (servUser *serviceUser) GetData(id int, name string) (*Domain, error) {
	fmt.Println("id service", id)
	result, err := servUser.repository.GetData(id, name)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}
