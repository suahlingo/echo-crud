package model

type CreateUserRequest struct {
	Name string `gorm:"name"`
	Age  int    `gorm:"default:0"`
}

type GetUserRequest struct {
	ID int `gorm:"primaryKey"`
}

type GetUserListRequest struct {
	Name string `query:"name"`
	Age  int    `query:"age"`
}

// 둘 중 한개만 수정 가능, 하나의 값만 입락하면 다른값은 기존 값으로 반환, 포인터타입 변수 사용
type UpdateUserRequest struct {
	ID   int     `gorm:"primaryKey"`
	Name *string `json:"name"`
	Age  *int    `json:"age"`
}

type DeleteUserRequest struct {
	ID int `gorm:"primaryKey"`
}
