package waste_test

import (
	"net/http"
	"os"
	"strings"
	"testing"
	_wasteHandler "wastebank-ca/app/presenter/waste"
	_wasteReq "wastebank-ca/app/presenter/waste/request"
	_wasteRes "wastebank-ca/app/presenter/waste/response"
	"wastebank-ca/bussines/waste"
	_wasteMocks "wastebank-ca/bussines/waste/mocks"

	"net/http/httptest"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	wasteServiceMock _wasteMocks.Service
	wasteHandler     _wasteHandler.Presenter
	// wasteRepository  _wasteMocks.Repository

	wasteDomain           waste.DomainWaste
	wasteDomainArray      []waste.DomainWaste
	wasteReqUpdate        _wasteReq.UpdateRequest
	wasteReqInsert        string
	wasteReqInsertInvalid string
	wasteResp             _wasteRes.Waste
	wasteRespCategory     _wasteRes.WasteCategory
)

func TestMain(m *testing.M) {
	wasteHandler = *_wasteHandler.NewHandler(&wasteServiceMock)

	wasteDomain = waste.DomainWaste{
		ID:            1,
		Name:          "paper",
		CategoryId:    1,
		TotalStock:    23,
		PurchasePrice: 1000,
	}

	wasteDomainArray = []waste.DomainWaste{
		{
			ID:            1,
			Name:          "paper",
			CategoryId:    1,
			TotalStock:    23,
			PurchasePrice: 1000,
		},
		{
			ID:            2,
			Name:          "paper2",
			CategoryId:    1,
			TotalStock:    23,
			PurchasePrice: 1000,
		},
	}

	wasteReqUpdate = _wasteReq.UpdateRequest{
		ID: 1,
		Data: waste.DomainWaste{
			ID:         1,
			Name:       "plastic",
			CategoryId: 1,
		},
	}

	wasteReqInsertInvalid = `{
		ID:         1,
		Name:       "waste test",
		CategoryId 1,
		PurchasePrice: 200,
		TotalStock: 102,
	}`

	wasteReqInsert = `{"id":1, "name":"waste test", "categoryId":1, "purchasePrice":200, "totalStock":102}`

	wasteResp = _wasteRes.Waste{
		ID:            1,
		Name:          "plastic",
		CategoryId:    1,
		PurchasePrice: 2000,
	}

	wasteRespCategory = _wasteRes.WasteCategory{
		ID:   1,
		Name: "organic",
	}

	os.Exit(m.Run())

}

func TestInsert(t *testing.T) {
	t.Run("Insert | Valid", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/waste", strings.NewReader(wasteReqInsert))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		wasteServiceMock.On("Append", mock.Anything).Return(&wasteDomain, nil).Once()
		if assert.NoError(t, wasteHandler.Insert(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("Insert | Invalid Bind", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/waste", strings.NewReader(wasteReqInsertInvalid))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		if assert.NoError(t, wasteHandler.Insert(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Insert | Invalid database", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/waste", strings.NewReader(wasteReqInsert))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		wasteServiceMock.On("Append", mock.Anything).Return(nil, assert.AnError).Once()
		if assert.NoError(t, wasteHandler.Insert(c)) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
		}
	})

}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/waste", strings.NewReader(wasteReqInsert))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		wasteServiceMock.On("Update", mock.Anything).Return(&wasteDomain, nil).Once()
		if assert.NoError(t, wasteHandler.Update(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("Update | Invalid Bind", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/waste", strings.NewReader(wasteReqInsertInvalid))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		if assert.NoError(t, wasteHandler.Update(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("Update | Invalid database", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/waste", strings.NewReader(wasteReqInsert))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		wasteServiceMock.On("Update", mock.Anything).Return(nil, assert.AnError).Once()
		if assert.NoError(t, wasteHandler.Update(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

}

func TestGetData(t *testing.T) {
	t.Run("GetData | Valid", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/waste", strings.NewReader(wasteReqInsert))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		// name := c.QueryParam("name")
		// id, _ := strconv.Atoi(c.QueryParam("id"))

		wasteServiceMock.On("GetData", mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(&wasteDomain, nil).Once()
		if assert.NoError(t, wasteHandler.GetData(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("GetData | Invalid database", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/waste", strings.NewReader(wasteReqInsert))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		wasteServiceMock.On("GetData", mock.AnythingOfType("int"), mock.AnythingOfType("string")).Return(nil, assert.AnError).Once()
		if assert.NoError(t, wasteHandler.GetData(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

}

func TestFindAll(t *testing.T) {
	t.Run("FindAll | Valid", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/waste", strings.NewReader(wasteReqInsert))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		wasteServiceMock.On("FindAll").Return(&wasteDomainArray, nil).Once()
		if assert.NoError(t, wasteHandler.FindAll(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("FindAll | Invalid", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/waste", strings.NewReader(wasteReqInsert))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		wasteServiceMock.On("FindAll").Return(nil, assert.AnError).Once()
		if assert.NoError(t, wasteHandler.FindAll(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

}
