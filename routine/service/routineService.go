package service

import (
	"awesomeProject/routine/model"
	"math/rand"
)

// 사용자가 새로고침 여부를 true로 보낼때 for문을 이용해서 저장해둔 20개의 랜덤 숫자를 한개씩 사용자에게 반환하기
type RoutineService struct {
	RefreshArr   model.RefreshArr
	CurrentIndex int
}

func NewRoutineService() *RoutineService {
	refreshArr := InitRandomArray() // 랜덤 배열 초기화
	return &RoutineService{
		RefreshArr:   refreshArr,
		CurrentIndex: 0,
	}
}

var currunteIndex int = 0

func InitRandomArray() model.RefreshArr {
	var refreshArr model.RefreshArr
	for i := 0; i < 20; i++ {
		refreshArr.Arr[i] = rand.Intn(100)
	}
	return refreshArr
}

func GetRandomVaue(req model.RefreshRequest, refreshArr *model.RefreshArr) *model.RefreshResponse {
	if req.IsRefreshed {
		if currunteIndex >= len(refreshArr.Arr) {
			currunteIndex = 0
		}

		resp := &model.RefreshResponse{
			Num:   refreshArr.Arr[currunteIndex],
			Index: currunteIndex,
		}

		currunteIndex++

		return resp
	}
	return nil

}
