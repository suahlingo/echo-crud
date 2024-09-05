package dto

type RefreshArr struct {
	ID  int   `gorm:"primaryKey;autoIncrement" json:"id"`
	Arr []int `gorm:"not null" json:"arr"`
}

type RefreshResponse struct {
	Arr []int `json:"arr"`
}
