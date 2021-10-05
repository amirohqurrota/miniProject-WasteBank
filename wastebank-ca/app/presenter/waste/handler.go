package waste

import (
	"fmt"
	"net/http"
	"strconv"
	"wastebank-ca/app/presenter/waste/request"
	"wastebank-ca/app/presenter/waste/response"
	"wastebank-ca/bussines/waste"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	serviceWaste waste.Service
}

func NewHandler(wasteServ waste.Service) *Presenter {
	return &Presenter{
		serviceWaste: wasteServ,
	}
}

func (handler *Presenter) Insert(echoContext echo.Context) error {
	var req request.WasteInsert
	if err := echoContext.Bind(&req); err != nil {
		return echoContext.JSON(http.StatusBadRequest, "something wrong")
	}

	domain := request.ToDomainInsert(req)
	resp, err := handler.serviceWaste.Append(domain)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, "something wrong in your request")
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
	//fmt.Println("id handler", domain.ID)
	resp, err := handler.serviceWaste.Update(domain)
	if err != nil {
		return echoContext.JSON(http.StatusBadRequest, "something wrong")
	}
	return echoContext.JSON(http.StatusOK, response.FromDomain(*resp))
}

func (handler *Presenter) FindAll(echoContext echo.Context) error {
	resp, err := handler.serviceWaste.FindAll()
	if err != nil {
		return echoContext.JSON(http.StatusBadRequest, "user not found")
	}
	fmt.Println("array cond in presenter ", len(*resp))

	wasteDomain := []response.Waste{}
	for _, value := range *resp {
		wasteDomain = append(wasteDomain, response.FromDomain(value))
	}

	return echoContext.JSON(http.StatusOK, wasteDomain)
}

func (handler *Presenter) GetData(echoContext echo.Context) error {
	//var req request.WasteInsert
	// if err := echoContext.Bind(&req); err != nil {
	// 	return echoContext.JSON(http.StatusBadRequest, "something wrong in your request")
	// }
	name := echoContext.QueryParam("name")
	id, _ := strconv.Atoi(echoContext.QueryParam("id"))

	resp, err := handler.serviceWaste.GetData(id, name)

	if err != nil {
		return echoContext.JSON(http.StatusBadRequest, "not found")
	}
	return echoContext.JSON(http.StatusOK, resp)
}
