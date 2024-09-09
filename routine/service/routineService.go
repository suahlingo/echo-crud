package service

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
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
		log.Printf("Error fetching data from %s: %v", s.BaseURL, err)
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	// 문자열 -> 숫자로 변환
	//dataStr := strings.TrimSpace(string(body))
	//number, err := strconv.Atoi(dataStr)
	//if err != nil {
	//	return 0, err
	//}

	var number int
	err = json.Unmarshal(body, &number)
	if err != nil {
		return 0, err
	}

	return number, nil
}
