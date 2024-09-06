package model

//func Migrate(db *gorm.DB) error {
//	return db.AutoMigrate(&User{})
//}

type User struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"not null" json:"name"`
	Age  int    `gorm:"not null" json:"age"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

type CreateUserResponse struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Message string `json:"message"`
}

type GetUserResponse struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type GetUserListResponse struct {
	Users []User `json:"users"`
}

type UpdateUserResponse struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Message string `json:"message"`
}
