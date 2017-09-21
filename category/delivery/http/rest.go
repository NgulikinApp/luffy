package http

import (
	"net/http"

	"strconv"

	"github.com/NgulikinApp/luffy/category/usecase"
	"github.com/labstack/echo"
)

type ResponseError struct {
	Message string `json:"message"`
}

type HTTPHandler struct {
	Usecase usecase.Usecase
}

func (h *HTTPHandler) Fetch(c echo.Context) error {
	num := 10
	if queryNum := c.QueryParam(`num`); queryNum != `` {
		var err error
		num, err = strconv.Atoi(queryNum)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &ResponseError{`Num parameter must be integer`})
		}
	}

	cursor := 0
	if queryCursor := c.QueryParam(`cursor`); queryCursor != `` {
		var err error
		cursor, err = strconv.Atoi(queryCursor)
		if err != nil {
			return c.JSON(http.StatusBadRequest, &ResponseError{`Cursor parameter must be integer`})
		}
	}

	res, err := h.Usecase.Fetch(int64(num), int64(cursor))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &ResponseError{err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func Init(e *echo.Echo, u usecase.Usecase) {
	handler := HTTPHandler{u}

	e.GET("/category", handler.Fetch)
}
