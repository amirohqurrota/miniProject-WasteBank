package transaction

import (
	"fmt"
	"net/http"
	"wastebank-ca/app/presenter/transaction/request"
	"wastebank-ca/app/presenter/transaction/response"
	"wastebank-ca/bussines/transaction"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	serviceTransaction transaction.Service
}

func NewHandler(transactionServ transaction.Service) *Presenter {
	return &Presenter{
		serviceTransaction: transactionServ,
	}
}

func (handler *Presenter) Insert(echoContext echo.Context) error {
	var req request.Transaction
	if err := echoContext.Bind(&req); err != nil {
		return echoContext.JSON(http.StatusBadRequest, "something wrong in you request")
	}

	domain := request.ToDomainTransaction(req)
	fmt.Println(domain.AdminID)
	resp, err := handler.serviceTransaction.Append(domain)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, "something wrong")
	}
	return echoContext.JSON(http.StatusOK, response.FromDomain(*resp))
}

func (handler *Presenter) AddNewType(echoContext echo.Context) error {
	var req request.TypeTransaction
	if err := echoContext.Bind(&req); err != nil {
		return echoContext.JSON(http.StatusBadRequest, "something wrong in your request")
	}
	domain := request.ToDomainTypeTrans(req)
	resp, err := handler.serviceTransaction.AddNewType(domain)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, "something wrong")
	}
	return echoContext.JSON(http.StatusOK, response.FromDomainType(*resp))

}

// func (handler *Presenter) Update(echoContext echo.Context) error {

// 	var req request.UpdateRequest
// 	if err := echoContext.Bind(&req); err != nil {
// 		return echoContext.JSON(http.StatusBadRequest, "something wrong")
// 	}
// 	fmt.Println("id handler", req.ID)
// 	domain := request.UpdateToDomain(req)
// 	fmt.Println("id handler", domain.ID)
// 	resp, err := handler.serviceUser.Update(domain)
// 	if err != nil {
// 		return echoContext.JSON(http.StatusBadRequest, "something wrong")
// 	}
// 	return echoContext.JSON(http.StatusOK, response.FromDomain(*resp))

// }

// func (handler *Presenter) FindByID(echoContext echo.Context) error {
// 	var req request.UserInsert
// 	if err := echoContext.Bind(&req); err != nil {
// 		return echoContext.JSON(http.StatusBadRequest, "something wrong")
// 	}
// 	domain := request.ToDomain(req)
// 	resp, err := handler.serviceUser.FindByID(domain.ID)
// 	fmt.Println("id handler", domain.ID)
// 	if err != nil {
// 		return echoContext.JSON(http.StatusBadRequest, "user not found")
// 	}
// 	return echoContext.JSON(http.StatusOK, response.FromDomain(*resp))

// }
