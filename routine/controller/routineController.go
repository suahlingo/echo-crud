package controller

import (
	"awesomeProject/routine/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

type routineController struct {
	RoutineService *service.RoutineService
	mutex          sync.Mutex
}

func NewRoutineController(routineService *service.RoutineService) *routineController {
	return &routineController{RoutineService: routineService}
}

func (r *routineController) GetDataList(c echo.Context) error {
	var randomData []int

	for len(randomData) < 20 {
		newData, err := r.RoutineService.FetchData()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch data"})
		}

		randomData = append(randomData, newData)
	}

	// 20개가 쌓이면 배열 반환
	return c.JSON(http.StatusOK, randomData)
}

func (r *routineController) GoRoutine(c echo.Context) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	var (
		results []int
		wg      sync.WaitGroup
		mu      sync.Mutex
	)

	wg.Add(20) //20개

	for i := 0; i < 20; i++ {
		go func() {
			defer wg.Done()
			newData, err := r.RoutineService.FetchData()
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch data"})
				return
			}
			mu.Lock()
			results = append(results, newData)
			mu.Unlock()
		}()
	}

	wg.Wait()
	return c.JSON(http.StatusOK, results)
}
