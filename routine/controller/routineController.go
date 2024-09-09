package controller

import (
	"awesomeProject/routine/service"
	"fmt"
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
		msgChan = make(chan string)
	)

	wg.Add(20) //20개

	//채널
	go func() {
		for msg := range msgChan {
			fmt.Println("Status:", msg)
		}
	}()

	for i := 0; i < 20; i++ {
		go func(i int) {
			defer wg.Done()

			//고루틴 시작할때 메시지
			msgChan <- fmt.Sprintf("Starting goroutine %d", i+1)

			newData, err := r.RoutineService.FetchData()
			if err != nil {
				c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch data"})
				return
			}
			mu.Lock()
			results = append(results, newData)
			mu.Unlock()

			msgChan <- fmt.Sprintf("Finished goroutine %d", i+1)
		}(i)
	}

	wg.Wait()
	close(msgChan)

	return c.JSON(http.StatusOK, results)
}
