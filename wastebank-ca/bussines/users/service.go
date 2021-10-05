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
	fmt.Println("id service", user.ID)
	result, err := servUser.repository.Update(user)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}
func (servUser *serviceUser) FindByID(id int) (*Domain, error) {
	fmt.Println("id service", id)
	result, err := servUser.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}
