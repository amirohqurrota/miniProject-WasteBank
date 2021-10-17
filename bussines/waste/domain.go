package waste

import "time"

type DomainCategory struct {
	ID   int
	Name string
}

type DomainWaste struct {
	ID            int
	Name          string
	CategoryId    int
	TotalStock    int
	PurchasePrice int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Service interface {
	Append(waste *DomainWaste) (*DomainWaste, error)
	Update(waste *DomainWaste) (*DomainWaste, error)
	FindAll() (*[]DomainWaste, error)
	GetData(id int, name string) (*DomainWaste, error)
}

type Repository interface {
	Insert(waste *DomainWaste) (*DomainWaste, error)
	Update(waste *DomainWaste) (*DomainWaste, error)
	FindAll() (*[]DomainWaste, error)
	GetData(id int, name string) (*DomainWaste, error)
}
