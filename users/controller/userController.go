package controller

import (
	"awesomeProject/users/model"
	"awesomeProject/users/service"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type userController struct {
	UserService *service.UserService
}

func NewUserController(userService *service.UserService) *userController {
	return &userController{UserService: userService}
}

func (uc *userController) AddUser(c echo.Context) error {
	name := c.FormValue("name")
	ageStr := c.FormValue("age")

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid age value"})
	}

	req := &model.CreateUserRequest{
		Name: name,
		Age:  age,
	}

	resp, err := uc.UserService.AddUser(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

func (uc *userController) GetUser(c echo.Context) error {
	idStr := c.QueryParam("id")
	if idStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID를 입력해주세요."})
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID는 정수여야 합니다."})
	}

	req := &model.GetUserRequest{
		ID: id,
	}

	resp, err := uc.UserService.GetUser(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "사용자 정보를 조회하는 데 실패했습니다."})
	}
	if resp == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "존재하지 않는 사용자입니다."})
	}

	return c.JSON(http.StatusOK, resp)
}

func (uc *userController) GetUserList(c echo.Context) error {
	var req model.GetUserListRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "잘못된 요청 형식입니다"})
	}

	log.Println("Received name:", req.Name)
	log.Println("Received age:", req.Age)

	if req.Name == "" && req.Age == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "검색 조건(이름 또는 나이)을 입력해주세요"})
	}

	// 유저 리스트 조회 요청
	resp, err := uc.UserService.GetUserList(&req)
	if err != nil {
		// 서비스에서 반환된 에러 처리
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

func (uc *userController) UpdateUser(c echo.Context) error {
	idStr := c.QueryParam("id")
	if idStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID를 입력해주세요."})
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID는 정수여야 합니다."})
	}

	name := c.FormValue("name")
	var namePtr *string
	if name != "" {
		namePtr = &name // 값이 있으면 포인터로 처리
	}

	ageStr := c.FormValue("age")
	var agePtr *int
	if ageStr != "" {
		ageValue, err := strconv.Atoi(ageStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "나이는 정수여야 합니다."})
		}
		agePtr = &ageValue
	}

	req := &model.UpdateUserRequest{
		ID:   id,
		Name: namePtr,
		Age:  agePtr,
	}

	log.Println("UpdateUserRequest:", req)

	// 서비스 레벨에서 사용자 정보 업데이트
	resp, err := uc.UserService.UpdateUser(req.ID, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}

func (uc *userController) DeleteUser(c echo.Context) error {
	idStr := c.FormValue("id")
	if idStr == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID를 입력해주세요."})
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "ID는 양의 정수여야 합니다."})
	}

	req := &model.DeleteUserRequest{
		ID: id,
	}

	err = uc.UserService.DeleteUser(req)
	if err != nil {
		if err.Error() == "사용자를 찾을 수 없습니다." {
			return c.JSON(http.StatusNotFound, map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "탈퇴 완료"})
}
