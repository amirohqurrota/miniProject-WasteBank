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
	UpdateSaldo(id int, saldo int) (*Domain, error)
	//UpdateWaste(id int, waste int) (*Domain, error)
	GetData(id int, name string) (*Domain, error)
}

type Repository interface {
	Insert(user *Domain) (*Domain, error)
	Update(user *Domain) (*Domain, error)
	UpdateSaldo(id int, saldo int) (*Domain, error)
	//UpdateWaste(id int, waste int) (*Domain, error)
	GetData(id int, name string) (*Domain, error)
}
