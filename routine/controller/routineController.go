package controller

import (
	"awesomeProject/routine/model"
	"awesomeProject/routine/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

var resfreshArr = service.InitRandomArray()

func randRefresh(c echo.Context) error {

	var req model.RefreshRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, "잚못된 요청")
	}

	resp := service.GetRandomVaue(req, &model.RefreshArr{})
	if resp != nil {
		return c.JSON(http.StatusOK, resp)
	}

	return c.JSON(http.StatusBadRequest, "false")
}
