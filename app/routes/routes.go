package routes

import (
	"wastebank-ca/app/presenter/admin"
	"wastebank-ca/app/presenter/transaction"
	"wastebank-ca/app/presenter/users"
	"wastebank-ca/app/presenter/waste"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerList struct {
	UserHandler        users.Presenter
	AdminHandler       admin.Presenter
	WasteHandler       waste.Presenter
	TransactionHandler transaction.Presenter
	JWTMiddleware      middleware.JWTConfig
}

func (handler *HandlerList) RouteRegister(e *echo.Echo) {
	users := e.Group("user")
	users.POST("/register", handler.UserHandler.Insert)
	users.GET("/login", handler.UserHandler.CreateToken)
	users.PUT("/update", handler.UserHandler.Update, middleware.JWTWithConfig(handler.JWTMiddleware))
	users.GET("/getBy", handler.UserHandler.GetData, middleware.JWTWithConfig(handler.JWTMiddleware))

	admin := e.Group("admin")
	admin.POST("/register", handler.AdminHandler.Insert)
	admin.GET("/login", handler.AdminHandler.CreateToken)
	admin.PUT("/update", handler.AdminHandler.Update, middleware.JWTWithConfig(handler.JWTMiddleware))
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
