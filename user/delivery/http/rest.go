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

type UserHTTPHandler struct {
	Usecase usecase.UserUsecase
}

func (self *UserHTTPHandler) GetByID(c echo.Context) error {
	queryID := c.Param(`id`)
	id, err := strconv.ParseInt(queryID, 10, 64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, &ResponseError{user.ErrIDParam.Error()})
	}

	res, err := self.Usecase.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &ResponseError{err.Error()})
	}

	if res == nil {
		return c.JSON(http.StatusNotFound, &ResponseError{user.ErrNotFound.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func Init(e *echo.Echo, u usecase.UserUsecase) {
	handler := UserHTTPHandler{u}

	e.GET("/user/:id", handler.GetByID)
}
