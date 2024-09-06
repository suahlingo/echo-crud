package model

type RefreshArr struct {
	ID           int     `gorm:"primaryKey;autoIncrement" json:"id"`
	Arr          [20]int `gorm:"-" json:"arr"`
	CurrentIndex int     `json:"-"`
}

type RefreshResponse struct {
	Num   int `json:"num"`
	Index int `json:"index"`
}
