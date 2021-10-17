package waste

import (
	_wasteDomain "wastebank-ca/bussines/waste"

	"gorm.io/gorm"
)

type repoWaste struct {
	DBConn *gorm.DB
}

func NewRepoMySQL(db *gorm.DB) _wasteDomain.Repository {
	return &repoWaste{
		DBConn: db,
	}
}

func (repo *repoWaste) Insert(waste *_wasteDomain.DomainWaste) (*_wasteDomain.DomainWaste, error) {
	recordWaste := fromDomain(*waste)

	if err := repo.DBConn.Save(&recordWaste).Error; err != nil {
		return &_wasteDomain.DomainWaste{}, err
	}
	result := toDomain(&recordWaste)
	return &result, nil
}

func (repo *repoWaste) Update(waste *_wasteDomain.DomainWaste) (*_wasteDomain.DomainWaste, error) {
	recordWaste := fromDomain(*waste)
	//fmt.Println("id sql 1", waste.ID)
	if err := repo.DBConn.Where("id=?", waste.ID).Updates(recordWaste).Error; err != nil {
		return &_wasteDomain.DomainWaste{}, err
	}
	//fmt.Println("id sql", recordWaste.ID)
	result := toDomain(&recordWaste)
	return &result, nil
}

func (repo *repoWaste) FindAll() (*[]_wasteDomain.DomainWaste, error) {
	recordWaste := []Waste{}
	if err := repo.DBConn.Find(&recordWaste).Error; err != nil {
		return &[]_wasteDomain.DomainWaste{}, err
	}

	wasteDomain := []_wasteDomain.DomainWaste{}
	for _, value := range recordWaste {
		wasteDomain = append(wasteDomain, toDomain(&value))
	}

	//fmt.Println(len(wasteDomain))

	return &wasteDomain, nil
}

func (repo *repoWaste) GetData(id int, name string) (*_wasteDomain.DomainWaste, error) {
	var recordWaste Waste
	if err := repo.DBConn.Where("name=? OR id=?", name, id).First(&recordWaste).Error; err != nil {
		return &_wasteDomain.DomainWaste{}, err
	}
	result := toDomain(&recordWaste)
	return &result, nil
}
