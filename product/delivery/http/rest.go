package http

import (
	"net/http"
	"strconv"

	"github.com/NgulikinApp/luffy/product"
	"github.com/NgulikinApp/luffy/product/usecase"
	"github.com/labstack/echo"
)

type ResponseError struct {
	Message string `json:"error"`
}

type HTTPHandler struct {
	Usecase usecase.Usecase
}

func (h HTTPHandler) Store(c echo.Context) error {
	p := product.Product{}
	c.Bind(&p)

	err := h.Usecase.Store(&p)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, &ResponseError{`Error`})
	}

	return c.JSON(http.StatusOK, p)
}

func (h HTTPHandler) Fetch(c echo.Context) error {
	filter := product.Filter{}

	num := int64(10)
	if queryNum := c.QueryParam(`num`); queryNum != `` {
		var err error
		num, err = strconv.ParseInt(queryNum, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &ResponseError{`Num parameter must be integer`})
		}
	}

	cursor := int64(0)
	if queryCursor := c.QueryParam(`cursor`); queryCursor != `` {
		var err error
		cursor, err = strconv.ParseInt(queryCursor, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &ResponseError{`Cursor parameter must be integer`})
		}
	}

	v, nc, err := h.Usecase.Fetch(filter, num, cursor)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &ResponseError{`Error`})
	}

	c.Response().Header().Set(`X-Cursor`, nc)
	return c.JSON(http.StatusOK, v)
}

func Init(e *echo.Echo, u usecase.Usecase) {
	handler := HTTPHandler{u}

	e.POST(`/product`, handler.Store)
	e.GET(`/product`, handler.Fetch)
}
