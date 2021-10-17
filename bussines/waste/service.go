package waste

type serviceWaste struct {
	repository Repository
}

func NewService(repoWaste Repository) Service {
	return &serviceWaste{
		repository: repoWaste,
	}
}

func (servWaste serviceWaste) Append(waste *DomainWaste) (*DomainWaste, error) {
	//response,err:=servWaste.Append()
	result, err := servWaste.repository.Insert(waste)
	if err != nil {
		return &DomainWaste{}, err
	}
	return result, nil
}

func (servWaste *serviceWaste) Update(waste *DomainWaste) (*DomainWaste, error) {
	//fmt.Println("id serviec", waste.ID)
	result, err := servWaste.repository.Update(waste)
	if err != nil {
		return &DomainWaste{}, err
	}
	return result, nil
}

func (servWaste *serviceWaste) FindAll() (*[]DomainWaste, error) {
	result, err := servWaste.repository.FindAll()
	if err != nil {
		return &[]DomainWaste{}, err
	}

	return result, nil
}

func (servWaste *serviceWaste) GetData(id int, name string) (*DomainWaste, error) {
	result, err := servWaste.repository.GetData(id, name)
	if err != nil {
		return &DomainWaste{}, err
	}
	return result, nil
}
