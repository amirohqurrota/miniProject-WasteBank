package users

import (
	"fmt"
	"net/http"
	"strconv"
	"wastebank-ca/app/presenter/users/request"
	"wastebank-ca/app/presenter/users/response"
	"wastebank-ca/bussines/users"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	serviceUser users.Service
}

func NewHandler(userServ users.Service) *Presenter {
	return &Presenter{
		serviceUser: userServ,
	}
}

func (handler *Presenter) Insert(echoContext echo.Context) error {
	var req request.UserInsert
	if err := echoContext.Bind(&req); err != nil {
		return echoContext.JSON(http.StatusBadRequest, "something wrong")
	}

	domain := request.ToDomain(req)
	resp, err := handler.serviceUser.Append(domain)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, "something wrong")
	}
	return echoContext.JSON(http.StatusOK, response.FromDomain(*resp))
}

func (handler *Presenter) Update(echoContext echo.Context) error {

	var req request.UpdateRequest
	if err := echoContext.Bind(&req); err != nil {
		return echoContext.JSON(http.StatusBadRequest, "something wrong")
	}
	fmt.Println("id handler", req.ID)
	domain := request.UpdateToDomain(req)
	fmt.Println("id handler", domain.ID)
	resp, err := handler.serviceUser.Update(domain)
	if err != nil {
		return echoContext.JSON(http.StatusBadRequest, "something wrong")
	}
	return echoContext.JSON(http.StatusOK, response.FromDomain(*resp))

}

func (handler *Presenter) GetData(echoContext echo.Context) error {
	//var req request.WasteInsert
	// if err := echoContext.Bind(&req); err != nil {
	// 	return echoContext.JSON(http.StatusBadRequest, "something wrong in your request")
	// }
	fmt.Println("handler user checking")
	firstName := echoContext.QueryParam("firstName")
	id, _ := strconv.Atoi(echoContext.QueryParam("id"))

	resp, err := handler.serviceUser.GetData(id, firstName)
	fmt.Println("handler user ", id)
	if err != nil {
		return echoContext.JSON(http.StatusBadRequest, "not found")
	}
	return echoContext.JSON(http.StatusOK, resp)
}
