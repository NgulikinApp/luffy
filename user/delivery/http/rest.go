package http

import (
	"net/http"

	"strconv"

	"github.com/NgulikinApp/luffy/user"
	"github.com/NgulikinApp/luffy/user/usecase"
	"github.com/labstack/echo"
)

type ResponseError struct {
	Message string `json:"message"`
}

type HTTPHandler struct {
	Usecase usecase.Usecase
}

func (h *HTTPHandler) GetByID(c echo.Context) error {
	queryID := c.Param(`id`)
	id, err := strconv.ParseInt(queryID, 10, 64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &ResponseError{user.ErrIDParam.Error()})
	}

	res, err := h.Usecase.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &ResponseError{err.Error()})
	}

	if res == nil {
		return c.JSON(http.StatusNotFound, &ResponseError{user.ErrNotFound.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func Init(e *echo.Echo, u usecase.Usecase) {
	handler := HTTPHandler{u}

	e.GET("/user/:id", handler.GetByID)
}
