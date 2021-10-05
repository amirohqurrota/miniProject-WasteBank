package users

import "time"

type Domain struct {
	ID         int
	FirstName  string
	LastName   string
	Telephone  string
	Address    string
	TotalWaste int
	TotalSaldo int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Service interface {
	Append(user *Domain) (*Domain, error)
	Update(user *Domain) (*Domain, error)
	FindByID(id int) (*Domain, error)
}

type Repository interface {
	Insert(user *Domain) (*Domain, error)
	Update(user *Domain) (*Domain, error)
	FindByID(id int) (*Domain, error)
}
