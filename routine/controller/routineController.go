package controller

import (
	"awesomeProject/routine/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type routineController struct {
	RoutineService *service.RoutineService
}

func NewRoutineController(routineService *service.RoutineService) *routineController {
	return &routineController{RoutineService: routineService}
}

func (r *routineController) GetDataList(c echo.Context) error {
	var randomData []int

	for len(randomData) < 20 {
		newData, err := r.RoutineService.FetchData() // 서비스 메서드 호출
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch data"})
		}

		randomData = append(randomData, newData)
	}

	// 20개가 쌓이면 배열 반환
	return c.JSON(http.StatusOK, randomData)
}

//type routineController struct {
//	RoutineService service.RoutineService
//}
//
//func NewRoutineController(routineService *service.Routin) *routineController {
//	return &routineController{RoutineService: *routineService}
//}

//func (rc *routineController) RandRefresh(c echo.Context) error {
//	// 쿼리 파라미터로 'isRefreshed' 값을 받음
//	isRefreshed := c.QueryParam("isRefreshed")
//
//	if isRefreshed == "true" {
//		resp := service.GetRandomValue(true, &rc.RoutineService.RefreshArr)
//		if resp != nil {
//			return c.JSON(http.StatusOK, resp)
//		}
//	}
//
//	return c.JSON(http.StatusBadRequest, "새로고침되지 않음")
//}

//func (rc *routineController) RandRefresh(c echo.Context) error {
//	var req model.RefreshRequest
//	if err := c.Bind(&req); err != nil {
//		return c.JSON(http.StatusBadRequest, "잘못된 요청")
//	}
//
//	resp := service.GetRandomValue(req, &rc.RoutineService.RefreshArr)
//	if resp != nil {
//		return c.JSON(http.StatusOK, resp)
//	}
//
//	return c.JSON(http.StatusBadRequest, "새로고침되지 않음")
//}
