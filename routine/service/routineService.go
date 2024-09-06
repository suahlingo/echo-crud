package service

import (
	"io"
	"net/http"
	"strconv"
	"strings"
)

type RoutineService struct {
	BaseURL string
}

func NewRoutineService(baseURL string) *RoutineService {
	return &RoutineService{
		BaseURL: baseURL,
	}
}

func (s *RoutineService) FetchData() (int, error) {
	resp, err := http.Get(s.BaseURL)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	// 문자열 -> 숫자로 변환
	dataStr := strings.TrimSpace(string(body))
	number, err := strconv.Atoi(dataStr)
	if err != nil {
		return 0, err
	}

	return number, nil
}

// 사용자가 새로고침 여부를 true로 보낼때 20개의 랜덤 숫자를 반환
//type RoutineService struct {
//	RefreshArr model.RefreshArr
//}
//
//func NewRoutineService() *RoutineService {
//	refreshArr := InitRandomArray()
//	return &RoutineService{
//		RefreshArr: refreshArr,
//	}
//}
//
//func InitRandomArray() model.RefreshArr {
//	var refreshArr model.RefreshArr
//	for i := 0; i < 20; i++ {
//		refreshArr.Arr[i] = rand.Intn(100) // 0~99
//	}
//	return refreshArr
//}
//
//func GetRandomValue(isRefreshed bool, refreshArr *model.RefreshArr) *model.RefreshResponse {
//	if isRefreshed {
//		if refreshArr.CurrentIndex >= len(refreshArr.Arr) {
//			refreshArr.CurrentIndex = 0
//		}
//
//		resp := &model.RefreshResponse{
//			Num:   refreshArr.Arr[refreshArr.CurrentIndex],
//			Index: refreshArr.CurrentIndex,
//		}
//
//		refreshArr.CurrentIndex++
//		return resp
//	}
//	return nil
//}
