package admin

import (
	"fmt"
	"net/http"
	"strconv"
	"wastebank-ca/app/presenter/admin/request"
	"wastebank-ca/app/presenter/admin/response"
	"wastebank-ca/bussines/admin"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	serviceAdmin admin.Service
}

func NewHandler(adminServ admin.Service) *Presenter {
	return &Presenter{
		serviceAdmin: adminServ,
	}
}

func (handler *Presenter) CreateToken(echoContext echo.Context) error {
	var req admin.Domain
	//ctx := echoContext.Request().Context()
	if err := echoContext.Bind(&req); err != nil {
		return echoContext.JSON(http.StatusBadRequest, "something wrong in your request")
	}
	username := req.Username
	password := req.Password
	token, err := handler.serviceAdmin.CreateToken(username, password)
	if err != nil {
		return err
	}
	req.Username = username
	return echoContext.JSON(http.StatusOK, response.FromDomainToken(token, req))
}

func (handler *Presenter) Insert(echoContext echo.Context) error {
	var req request.Admin
	if err := echoContext.Bind(&req); err != nil {
		return echoContext.JSON(http.StatusBadRequest, "something wrong")
	}

	domain := request.ToDomain(req)
	resp, err := handler.serviceAdmin.Append(domain)
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
	//fmt.Println("id handler", req.ID)
	domain := request.UpdateToDomain(req)
	resp, err := handler.serviceAdmin.Update(domain)
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
	username := echoContext.QueryParam("username")
	firstName := echoContext.QueryParam("firstName")
	lastName := echoContext.QueryParam("lastName")
	id, _ := strconv.Atoi(echoContext.QueryParam("id"))

	resp, err := handler.serviceAdmin.GetData(id, firstName, lastName, username)
	fmt.Println("handler user ", id)
	if err != nil {
		return echoContext.JSON(http.StatusBadRequest, "not found")
	}
	return echoContext.JSON(http.StatusOK, resp)
}
