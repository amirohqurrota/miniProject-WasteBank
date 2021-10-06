package routes

import (
	"wastebank-ca/app/presenter/admin"
	"wastebank-ca/app/presenter/transaction"
	"wastebank-ca/app/presenter/users"
	"wastebank-ca/app/presenter/waste"

	"github.com/labstack/echo/v4"
)

type HandlerList struct {
	UserHandler        users.Presenter
	AdminHandler       admin.Presenter
	WasteHandler       waste.Presenter
	TransactionHandler transaction.Presenter
}

func (handler *HandlerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("/register", handler.UserHandler.Insert)
	users.PUT("/update", handler.UserHandler.Update)
	users.GET("/getBy", handler.UserHandler.GetData)

	admin := e.Group("admin")
	admin.POST("/register", handler.AdminHandler.Insert)
	admin.PUT("/update", handler.AdminHandler.Update)
	admin.GET("/getBy", handler.AdminHandler.GetData)

	waste := e.Group("waste")
	waste.POST("/addNew", handler.WasteHandler.Insert)
	waste.PUT("/update", handler.WasteHandler.Update)
	waste.GET("/getAll", handler.WasteHandler.FindAll)
	waste.GET("/getBy", handler.WasteHandler.GetData)

	transaction := e.Group("transaction")
	transaction.POST("/add", handler.TransactionHandler.Insert)
	transaction.POST("/newType", handler.TransactionHandler.AddNewType)

}
