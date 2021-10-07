package main

import (
	"log"
	//"wastebank-ca/app/routes"
	handlerUsers "wastebank-ca/app/presenter/users"
	routes "wastebank-ca/app/routes"
	"wastebank-ca/bussines/transactions"
	"wastebank-ca/bussines/users"
	repoUsers "wastebank-ca/repository/sql/users"

	handlerAdmin "wastebank-ca/app/presenter/admin"
	"wastebank-ca/bussines/admin"
	repoAdmin "wastebank-ca/repository/sql/admin"

	handlerWaste "wastebank-ca/app/presenter/waste"
	"wastebank-ca/bussines/waste"
	repoWaste "wastebank-ca/repository/sql/waste"

	handlerTransaction "wastebank-ca/app/presenter/transaction"
	repoTransaction "wastebank-ca/repository/sql/transactions"

	"time"
	_middleware "wastebank-ca/app/middleware"

	repoNewsApi "wastebank-ca/repository/newsApi"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connectionString := "root:22juli1998@tcp(0.0.0.0:3306)/wastebank_miniproject?parseTime=True"

	var err error
	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(
		&repoUsers.User{},
		&repoAdmin.Admin{},
		&repoWaste.Waste{},
		&repoWaste.WasteCategory{},
		&repoTransaction.Transaction{},
		&repoTransaction.WasteDeposit{},
		&repoTransaction.TransactionType{},
	)

	return DB
}

func main() {
	db := initDB()
	e := echo.New()
	configJWT := _middleware.ConfigJWT{
		SecretJWT:       "secret!@#$%",
		ExpiresDuration: 1,
	}

	timeoutContext := time.Duration(2 * time.Second)

	// factory of domain
	usersRepo := repoUsers.NewRepoMySQL(db)
	usersServ := users.NewService(usersRepo, &configJWT, timeoutContext)
	usersHandler := handlerUsers.NewHandler(usersServ)

	adminRepo := repoAdmin.NewRepoMySQL(db)
	adminServ := admin.NewService(adminRepo, &configJWT, timeoutContext)
	adminHandler := handlerAdmin.NewHandler(adminServ)

	wasteRepo := repoWaste.NewRepoMySQL(db)
	wasteServ := waste.NewService(wasteRepo)
	wasteHandler := handlerWaste.NewHandler(wasteServ)

	newsApiRepo := repoNewsApi.NewNewsApi()

	transactionRepo := repoTransaction.NewRepoMySQL(db)
	transactionServ := transactions.NewService(transactionRepo, adminServ, usersServ, newsApiRepo)
	transactionHandler := handlerTransaction.NewHandler(transactionServ)

	// initial of routes
	routesInit := routes.HandlerList{
		UserHandler:        *usersHandler,
		WasteHandler:       *wasteHandler,
		TransactionHandler: *transactionHandler,
		AdminHandler:       *adminHandler,
		JWTMiddleware:      configJWT.Init(),
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(":8080"))
}
