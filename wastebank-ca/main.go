package main

import (
	"log"
	//"wastebank-ca/app/routes"
	handlerUsers "wastebank-ca/app/presenter/users"
	routes "wastebank-ca/app/routes"
	"wastebank-ca/bussines/users"
	repoUsers "wastebank-ca/repository/sql/users"

	handlerWaste "wastebank-ca/app/presenter/waste"
	"wastebank-ca/bussines/waste"
	repoWaste "wastebank-ca/repository/sql/waste"

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
		&repoWaste.Waste{},
		&repoWaste.WasteCategory{},
	)

	return DB
}

func main() {
	db := initDB()
	e := echo.New()

	// factory of domain
	usersRepo := repoUsers.NewRepoMySQL(db)
	usersServ := users.NewService(usersRepo)
	usersHandler := handlerUsers.NewHandler(usersServ)

	wasteRepo := repoWaste.NewRepoMySQL(db)
	wasteServ := waste.NewService(wasteRepo)
	wasteHandler := handlerWaste.NewHandler(wasteServ)

	// initial of routes
	routesInit := routes.HandlerList{
		UserHandler:  *usersHandler,
		WasteHandler: *wasteHandler,
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(":8080"))
}
