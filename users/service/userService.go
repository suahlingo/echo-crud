package service

import (
	"awesomeProject/users/model"
	"errors"
	"gorm.io/gorm"
)

// 서비스 구조체 정의
type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) AddUser(req *model.CreateUserRequest) (*model.CreateUserResponse, error) {

	if req.Name == "" {
		return nil, errors.New("이름은 필수 입력 사항입니다.")
	}
	if req.Age <= 0 {
		return nil, errors.New("나이는 0보다 커야 합니다.")
	}

	user := model.User{
		Name: req.Name,
		Age:  req.Age,
	}

	if err := s.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	resp := &model.CreateUserResponse{
		Name:    user.Name,
		Age:     user.Age,
		Message: "회원가입 완료",
	}

	return resp, nil
}

func (s *UserService) GetUser(req *model.GetUserRequest) (*model.GetUserResponse, error) {
	var user model.User

	if err := s.DB.Where("id = ?", req.ID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	resp := &model.GetUserResponse{
		Name: user.Name,
		Age:  user.Age,
	}

	return resp, nil
}

func (s *UserService) GetUserList(req *model.GetUserListRequest) (*model.GetUserListResponse, error) {
	var users []model.User
	query := s.DB.Model(&model.User{})

	// 이름이 있는 경우 이름으로 필터링
	if req.Name != "" {
		query = query.Where("name = ?", req.Name)
	}

	// 나이가 있는 경우 나이로 필터링
	if req.Age != 0 {
		query = query.Where("age = ?", req.Age)
	}

	result := query.Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	if len(users) == 0 {
		return nil, errors.New("해당하는 사용자 정보가 없습니다")
	}

	resp := &model.GetUserListResponse{
		Users: users,
	}

	return resp, nil
}

func (s *UserService) UpdateUser(id int, req *model.UpdateUserRequest) (*model.UpdateUserResponse, error) {
	var user model.User

	// ID로 사용자 조회
	if err := s.DB.Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("사용자를 찾을 수 없습니다.")
		}
		return nil, err
	}

	// 수정할 필드 업데이트
	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Age != nil {
		user.Age = *req.Age
	}

	// 수정된 사용자 정보 저장
	if err := s.DB.Save(&user).Error; err != nil {
		return nil, err
	}

	// 성공 메시지와 함께 응답 생성
	resp := &model.UpdateUserResponse{
		Name:    user.Name,
		Age:     user.Age,
		Message: "회원정보 수정 완료",
	}

	return resp, nil
}

func (s *UserService) DeleteUser(req *model.DeleteUserRequest) error {
	result := s.DB.Delete(&model.User{}, req.ID)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("사용자를 찾을 수 없습니다.")
	}

	return nil
}
