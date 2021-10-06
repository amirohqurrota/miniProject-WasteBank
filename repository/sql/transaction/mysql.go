package transaction

import (
	"fmt"
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
	recWasteDeposit := fromDomainDeposit(*wasteDeposit)

	if err := repo.DBConn.Save(&recWasteDeposit).Error; err != nil {
		return &_TransactionDomain.DomainDeposit{}, err
	}
	fmt.Println("sql new deposit aman")
	result := ToDomainDeposit(&recWasteDeposit)
	return &result, nil
}

func (repo *repoTransaction) Insert(trans *_TransactionDomain.DomainTransaction) (*_TransactionDomain.DomainTransaction, error) {
	//var idDeposit int
	if trans.TypeID == 1 {
		for n := range trans.DepositData {
			data, _ := repo.NewDeposit(&trans.DepositData[n])
			trans.DepositData[n].ID = data.ID
			trans.DepositData[n].TotalHeight = data.TotalHeight
			trans.DepositData[n].ID = data.ID
			trans.DepositData[n].WasteId = data.WasteId
		}
	}

	fmt.Println("--", trans.DepositData[0].TotalHeight)
	recordTransaction := fromDomainTrans(*trans)
	//fmt.Println(recordTransaction.AdminID, "in repository")
	if err := repo.DBConn.Save(&recordTransaction).Error; err != nil {
		return &_TransactionDomain.DomainTransaction{}, err
	}
	result := toDomainTrans(&recordTransaction)
	return &result, nil
}

func (repo *repoTransaction) AddNewType(typeTransaction *_TransactionDomain.DomainType) (*_TransactionDomain.DomainType, error) {
	recType := fromDomainType(*typeTransaction)

	if err := repo.DBConn.Save(&recType).Error; err != nil {
		return &_TransactionDomain.DomainType{}, err
	}
	result := toDomainType(&recType)
	//fmt.Println(result.Name)
	return &result, nil
}

// func (repo *repoTransaction) FindAll() (*[]_TransactionDomain.DomainTransaction, error) {
// 	// 	recordWaste := []Waste{}
// 	// 	if err := repo.DBConn.Find(&recordWaste).Error; err != nil {
// 	// 		return &[]_TransactionDomain.DomainTransaction{}, err
// 	// 	}

// 	// 	wasteDomain := []_TransactionDomain.DomainTransaction{}
// 	// 	for _, value := range recordWaste {
// 	// 		wasteDomain = append(wasteDomain, toDomain(&value))
// 	//}

// 	// 	//fmt.Println(len(wasteDomain))

// 	// 	return &wasteDomain, nil
// 	return nil, nil
// }

// func (repo *repoTransaction) GetData(id int, name string) (*_TransactionDomain.DomainTransaction, error) {
// 	// 	var recordWaste Waste
// 	// 	if err := repo.DBConn.Where("name=? OR id=?", name, id).First(&recordWaste).Error; err != nil {
// 	// 		return &_TransactionDomain.Domain{}, err
// 	// 	}
// 	// 	result := toDomain(&recordWaste)
// 	// 	return &result, nil
// 	return nil, nil
// }
