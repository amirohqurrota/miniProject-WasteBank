package admin

import "time"

type Domain struct {
	ID         int
	Username   string
	Password   string
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
	GetData(id int, firstName string, lastName string, username string) (*Domain, error)
	CreateToken(username, password string) (string, error)
}

type Repository interface {
	Insert(admin *Domain) (*Domain, error)
	Update(admin *Domain) (*Domain, error)
	GetData(id int, firstName string, lastName string, username string) (*Domain, error)
}
