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

func (handler *Presenter) CreateToken(echoContext echo.Context) error {
	var req users.Domain
	if err := echoContext.Bind(&req); err != nil {
		return echoContext.JSON(http.StatusBadRequest, "something wrong in your request")
	}
	username := req.Username
	password := req.Password
	token, err := handler.serviceUser.CreateToken(username, password)
	if err != nil {
		return err
	}
	//a, _ := auth.ParsingToken(token)
	req.Username = username
	return echoContext.JSON(http.StatusOK, response.FromDomainToken(token, req))
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
	//fmt.Println("handler user checking")
	firstName := echoContext.QueryParam("firstName")
	lastName := echoContext.QueryParam("lastName")
	username := echoContext.QueryParam("username")
	id, _ := strconv.Atoi(echoContext.QueryParam("id"))

	// token := echoContext.Get("admin").(*jwt.Token).Raw
	// if middleware.RoleValidation(token, "admin") {
	// 	resp, err := handler.serviceUser.GetData(id, firstName, lastName, username)
	// 	if err != nil {
	// 		return echoContext.JSON(http.StatusBadRequest, "not found")
	// 	}
	// 	return echoContext.JSON(http.StatusOK, resp)
	// }
	// return errors.New("forbidden role")

	resp, err := handler.serviceUser.GetData(id, firstName, lastName, username)
	if err != nil {
		return echoContext.JSON(http.StatusBadRequest, "not found")
	}
	return echoContext.JSON(http.StatusOK, resp)
}
