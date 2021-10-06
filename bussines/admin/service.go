package admin

import "fmt"

type serviceAdmin struct {
	repository Repository
}

func NewService(repoAdmin Repository) Service {
	return &serviceAdmin{
		repository: repoAdmin,
	}
}

func (servAdmin serviceAdmin) Append(admin *Domain) (*Domain, error) {
	//response,err:=servUser.Append()
	result, err := servAdmin.repository.Insert(admin)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (servAdmin *serviceAdmin) Update(admin *Domain) (*Domain, error) {
	fmt.Println("id service", admin.ID)
	result, err := servAdmin.repository.Update(admin)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}
func (servAdmin *serviceAdmin) GetData(id int, name string) (*Domain, error) {
	fmt.Println("id service", id)
	result, err := servAdmin.repository.GetData(id, name)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}
