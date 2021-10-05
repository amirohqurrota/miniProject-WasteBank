package transaction

import (
	_TransactionDomain "wastebank-ca/bussines/transaction"

	"gorm.io/gorm"
)

type repoTransaction struct {
	DBConn *gorm.DB
}

func NewRepoMySQL(db *gorm.DB) _TransactionDomain.Repository {
	return &repoTransaction{
		DBConn: db,
	}
}

func (repo *repoTransaction) NewDeposit(wasteDeposit *_TransactionDomain.DomainDeposit) (*_TransactionDomain.DomainDeposit, error) {
	recWasteDeposit := wasteDeposit

	if err := repo.DBConn.Save(&recWasteDeposit).Error; err != nil {
		return &_TransactionDomain.DomainDeposit{}, err
	}
	return recWasteDeposit, nil
}

func (repo *repoTransaction) Insert(trans *_TransactionDomain.DomainTransaction) (*_TransactionDomain.DomainTransaction, error) {
	var idDeposit int
	if trans.TypeID == 1 {
		data, _ := repo.NewDeposit(&trans.DepositData)
		idDeposit = data.ID
	}
	trans.DepositID = idDeposit
	recordTransation := fromDomain(*trans)

	if err := repo.DBConn.Save(&recordTransation).Error; err != nil {
		return &_TransactionDomain.DomainTransaction{}, err
	}
	result := toDomain(&recordTransation)
	return &result, nil
}

func (repo *repoTransaction) Update(waste *_TransactionDomain.DomainTransaction) (*_TransactionDomain.DomainTransaction, error) {
	// 	recordWaste := fromDomain(*waste)
	// 	//fmt.Println("id sql 1", waste.ID)
	// 	if err := repo.DBConn.Where("Id=?", waste.ID).Updates(recordWaste).Error; err != nil {
	// 		return &_TransactionDomain.Domain{}, err
	// 	}
	// 	//fmt.Println("id sql", recordWaste.ID)
	// 	result := toDomain(&recordWaste)
	// 	return &result, nil
	return nil, nil
}

func (repo *repoTransaction) FindAll() (*[]_TransactionDomain.DomainTransaction, error) {
	// 	recordWaste := []Waste{}
	// 	if err := repo.DBConn.Find(&recordWaste).Error; err != nil {
	// 		return &[]_TransactionDomain.DomainTransaction{}, err
	// 	}

	// 	wasteDomain := []_TransactionDomain.DomainTransaction{}
	// 	for _, value := range recordWaste {
	// 		wasteDomain = append(wasteDomain, toDomain(&value))
	//}

	// 	//fmt.Println(len(wasteDomain))

	// 	return &wasteDomain, nil
	return nil, nil
}

func (repo *repoTransaction) GetData(id int, name string) (*_TransactionDomain.DomainTransaction, error) {
	// 	var recordWaste Waste
	// 	if err := repo.DBConn.Where("name=? OR id=?", name, id).First(&recordWaste).Error; err != nil {
	// 		return &_TransactionDomain.Domain{}, err
	// 	}
	// 	result := toDomain(&recordWaste)
	// 	return &result, nil
	return nil, nil
}
