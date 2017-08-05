package http_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NgulikinApp/luffy/user"
	handler "github.com/NgulikinApp/luffy/user/delivery/http"
	"github.com/NgulikinApp/luffy/user/usecase/mocks"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userData = user.User{
		Username:  `gama`,
		Fullname:  `andhika gama`,
		DOB:       `1992-03-23`,
		Gender:    `male`,
		Source:    `web`,
		Activated: true,
	}
)

func TestHandlerGetByID(t *testing.T) {
	mockUser := userData
	mockUserUcase := new(mocks.UserUsecase)
	mockUserUcase.On("GetByID",
		mock.AnythingOfType("int64"),
	).Return(&mockUser, nil).Once()

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetPath(`user/:id`)
	c.SetParamNames(`id`)
	c.SetParamValues(`1`)

	h := new(handler.UserHTTPHandler)
	h.Usecase = mockUserUcase

	if assert.NoError(t, h.GetByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
	mockUserUcase.AssertExpectations(t)
}

func TestHandlerGetByIDNotFound(t *testing.T) {
	mockUserUcase := new(mocks.UserUsecase)
	mockUserUcase.On("GetByID",
		mock.AnythingOfType("int64"),
	).Return(nil, nil).Once()

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetPath(`user/:id`)
	c.SetParamNames(`id`)
	c.SetParamValues(`10`)

	h := new(handler.UserHTTPHandler)
	h.Usecase = mockUserUcase

	if assert.NoError(t, h.GetByID(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
	mockUserUcase.AssertExpectations(t)
}

func TestHandlerGetByIDBadParam(t *testing.T) {
	mockUserUcase := new(mocks.UserUsecase)
	mockUserUcase.On("GetByID",
		mock.AnythingOfType("int64"),
	).Return(nil, nil)

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetPath(`user/:id`)
	c.SetParamNames(`id`)
	c.SetParamValues(`a`)

	h := new(handler.UserHTTPHandler)
	h.Usecase = mockUserUcase

	if assert.NoError(t, h.GetByID(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
	mockUserUcase.AssertNotCalled(t, `GetByID`, int64(1))
}

func TestHandlerGetByIDError(t *testing.T) {
	mockUserUcase := new(mocks.UserUsecase)
	mockUserUcase.On("GetByID",
		mock.AnythingOfType("int64"),
	).Return(nil, errors.New(`Error`)).Once()

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetPath(`user/:id`)
	c.SetParamNames(`id`)
	c.SetParamValues(`1`)

	h := new(handler.UserHTTPHandler)
	h.Usecase = mockUserUcase

	if assert.NoError(t, h.GetByID(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	}
	mockUserUcase.AssertExpectations(t)
}

/*
// Test Handler Get Company Data By ID Not Found
func TestHandlerGetByIDNotFound(t *testing.T) {
	mockCompanyUsecase := new(mocks.CompanyUsecase)
	mockCompanyUsecase.On("GetByID", mock.AnythingOfType("int64")).Return(nil, nil)

	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetPath(`company/:id`)
	c.SetParamNames(`id`)
	c.SetParamValues(`1000`)

	h := new(handler.CompanyHTTPHandler)
	h.Usecase = mockCompanyUsecase

	if assert.NoError(t, h.GetByID(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}

// Test Handler Get Company Data By ID Bad Param
func TestHandlerGetByIDBadParam(t *testing.T) {
	mockCompanyUsecase := new(mocks.CompanyUsecase)
	mockCompanyUsecase.On("GetByID", mock.AnythingOfType("int64"),
		mock.AnythingOfType("string")).Return(http.StatusBadRequest, errors.New(`Error`))

	q := make(url.Values)
	q.Set("id", "1")

	req := httptest.NewRequest(echo.GET, "/company?"+q.Encode(), nil)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)

	h := new(handler.CompanyHTTPHandler)
	h.Usecase = mockCompanyUsecase

	if assert.NoError(t, h.GetByID(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

// Test Handler Get Company Data By ID Error
func TestHandlerGetByIDError(t *testing.T) {
	mockCompanyUsecase := new(mocks.CompanyUsecase)
	mockCompanyUsecase.On("GetByID", mock.AnythingOfType("int64")).Return(nil, errors.New(`error`))

	req := httptest.NewRequest(echo.GET, "/?", nil)
	rec := httptest.NewRecorder()

	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetPath(`company/:id`)
	c.SetParamNames(`id`)
	c.SetParamValues(`1`)

	h := new(handler.CompanyHTTPHandler)
	h.Usecase = mockCompanyUsecase
	id := int64(1)

	if assert.NoError(t, h.GetByID(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	}
	mockCompanyUsecase.AssertCalled(t, `GetByID`, id)
}

// Test Handler Update
func TestHandlerUpdate(t *testing.T) {
	mockCompanyUsecase := new(mocks.CompanyUsecase)
	mockCompanyUsecase.On("Update", mock.AnythingOfType("*company.Company")).Return(nil)

	req := httptest.NewRequest(echo.PUT, "/company", nil)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetPath(`company/:id`)
	c.SetParamNames(`id`)
	c.SetParamValues(`1`)

	h := new(handler.CompanyHTTPHandler)
	h.Usecase = mockCompanyUsecase

	if assert.NoError(t, h.Update(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}*/
