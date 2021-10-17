package waste_test

import (
	"testing"
	"time"
	"wastebank-ca/bussines/waste"
	_wasteRepo "wastebank-ca/repository/sql/waste"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var listOfWaste = []waste.DomainWaste{
	{
		ID:            1,
		Name:          "paper",
		CategoryId:    1,
		TotalStock:    23,
		PurchasePrice: 1000,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	},
	{
		ID:            2,
		Name:          "paper2",
		CategoryId:    1,
		TotalStock:    23,
		PurchasePrice: 1000,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	},
}

func TestInsert(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	wasteRepo := _wasteRepo.NewRepoMySQL(gdb)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `wastes` SET `name`=?,`category_id`=?,`purchase_price`=?,`total_stock`=?,`created_at`=?,`updated_at`=? WHERE `id` = ?").
		WithArgs(listOfWaste[0].Name, listOfWaste[0].CategoryId, int64(listOfWaste[0].TotalStock), int64(listOfWaste[0].PurchasePrice), listOfWaste[0].CreatedAt, listOfWaste[0].UpdatedAt, listOfWaste[0].ID).
		WillReturnResult(sqlmock.NewResult(int64(1), int64(1)))
	mock.ExpectCommit()

	_, err = wasteRepo.Insert(&listOfWaste[0])
	require.Error(t, err)
	//require.NotEmpty(t, err)
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	wasteRepo := _wasteRepo.NewRepoMySQL(gdb)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `wastes` SET `name`=?,`category_id`=?,`purchase_price`=?,`total_stock`=?,`created_at`=?,`updated_at`=? WHERE `id` = ?").
		WithArgs(listOfWaste[0].Name, listOfWaste[0].CategoryId, int64(listOfWaste[0].TotalStock), int64(listOfWaste[0].PurchasePrice), listOfWaste[0].CreatedAt, listOfWaste[0].UpdatedAt, listOfWaste[0].ID).
		WillReturnResult(sqlmock.NewResult(int64(1), int64(1)))
	mock.ExpectCommit()

	_, err = wasteRepo.Update(&listOfWaste[0])
	require.Error(t, err)
}

func TestFindAll(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	wasteRepo := _wasteRepo.NewRepoMySQL(gdb)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("SELECT * FROM `wastes`")
	mock.ExpectCommit()

	_, err = wasteRepo.FindAll()
	require.Error(t, err)
	//require.NotEmpty(t, err)
}

func TestGetData(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	wasteRepo := _wasteRepo.NewRepoMySQL(gdb)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("SELECT * FROM `wastes`")
	mock.ExpectCommit()

	_, err = wasteRepo.GetData(1, "plastic")
	require.Error(t, err)
	//require.NotEmpty(t, err)
}

// func TestGetDataValid(t *testing.T) {
// 	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	gdb, _ := gorm.Open(mysql.New(mysql.Config{
// 		Conn:                      db,
// 		SkipInitializeWithVersion: true,
// 	}), &gorm.Config{})

// 	wasteRepo := _wasteRepo.NewRepoMySQL(gdb)
// 	defer db.Close()

// 	mock.ExpectBegin()
// 	mock.ExpectExec("SELECT * FROM `wastes` WHERE `name`=? OR `id` = ? ORDER BY `wastes`.`id` LIMIT 1").
// 		WithArgs(listOfWaste[0].ID, listOfWaste[0].Name).
// 		WillReturnResult(sqlmock.NewResult(int64(1), int64(0)))
// 	mock.ExpectCommit()

// 	_, err = wasteRepo.GetData(listOfWaste[0].ID, listOfWaste[0].Name)
// 	require.NoError(t, err)
// 	//require.NotEmpty(t, err)
// }
