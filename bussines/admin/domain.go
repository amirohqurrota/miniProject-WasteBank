package admin

import "time"

type Domain struct {
	ID         int
	FirstName  string
	LastName   string
	Telephone  string
	Address    string
	TotalBonus int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Service interface {
	Append(admin *Domain) (*Domain, error)
	Update(admin *Domain) (*Domain, error)
	GetData(id int, name string) (*Domain, error)
}

type Repository interface {
	Insert(admin *Domain) (*Domain, error)
	Update(admin *Domain) (*Domain, error)
	GetData(id int, name string) (*Domain, error)
}
