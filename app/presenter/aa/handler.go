package transaction

import (
	"net/http"
	"wastebank-ca/app/presenter/transaction/request"
	"wastebank-ca/app/presenter/transaction/response"
	_transactionService "wastebank-ca/bussines/transaction"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	serviceTransaction _transactionService.Service
}

func NewHandler(transactionServ _transactionService.Service) *Presenter {
	return &Presenter{
		serviceTransaction: transactionServ,
	}
}

func (handler *Presenter) Insert(echoContext echo.Context) error {
	var req request.Transaction
	if err := echoContext.Bind(&req); err != nil {
		return echoContext.JSON(http.StatusBadRequest, "something wrong  in your request")
	}

	domain := request.ToDomainTransaction(req)
	resp, err := handler.serviceTransaction.Append(domain)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, "something wrong")
	}
	return echoContext.JSON(http.StatusOK, response.FromDomain(*resp))
}

// func (handler *Presenter) Update(echoContext echo.Context) error {

// 	var req request.UpdateWasteDepo
// 	if err := echoContext.Bind(&req); err != nil {
// 		return echoContext.JSON(http.StatusBadRequest, "something wrong")
// 	}
// 	//fmt.Println("id handler", req.ID)
// 	domain := request.UpdateToDomain(req)
// 	//fmt.Println("id handler", domain.ID)
// 	resp, err := handler.serviceWasteDepo.Update(domain)
// 	if err != nil {
// 		return echoContext.JSON(http.StatusBadRequest, "something wrong")
// 	}
// 	return echoContext.JSON(http.StatusOK, response.FromDomain(*resp))
// }

// func (handler *Presenter) FindAll(echoContext echo.Context) error {
// 	resp, err := handler.serviceWasteDepo.FindAll()
// 	if err != nil {
// 		return echoContext.JSON(http.StatusBadRequest, "user not found")
// 	}
// 	fmt.Println("array cond in presenter ", len(*resp))

// 	wasteDomain := []response.WasteDeposit{}
// 	for _, value := range *resp {
// 		wasteDomain = append(wasteDomain, response.FromDomain(value))
// 	}

// 	return echoContext.JSON(http.StatusOK, wasteDomain)
// }

// func (handler *Presenter) GetData(echoContext echo.Context) error {
// 	//var req request.WasteInsert
// 	// if err := echoContext.Bind(&req); err != nil {
// 	// 	return echoContext.JSON(http.StatusBadRequest, "something wrong in your request")
// 	// }
// 	id, _ := strconv.Atoi(echoContext.QueryParam("id"))

// 	resp, err := handler.serviceWasteDepo.GetData(id)

// 	if err != nil {
// 		return echoContext.JSON(http.StatusBadRequest, "not found")
// 	}
// 	return echoContext.JSON(http.StatusOK, resp)
// }
