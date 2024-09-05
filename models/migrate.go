package models

import (
	"encoding/json"
	"gorm.io/gorm"
)

type User struct {
	ID    int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"not null" json:"name"`
	Email string `gorm:"not null;unique" json:"email"`
	Age   int    `gorm:"not null" json:"age"`
}

// RefreshArr 구조체 (배열을 JSON으로 저장)
type RefreshArr struct {
	ID      int     `gorm:"primaryKey;autoIncrement" json:"id"`
	Arr     [20]int `gorm:"-" json:"arr"` // 배열 필드
	ArrJSON []byte  `gorm:"type:json" json:"-"`
}

// BeforeCreate: 배열을 JSON으로 변환해 저장
func (r *RefreshArr) BeforeCreate(tx *gorm.DB) (err error) {
	r.ArrJSON, err = json.Marshal(r.Arr)
	return
}

// AfterFind: JSON을 다시 배열로 변환
func (r *RefreshArr) AfterFind(tx *gorm.DB) (err error) {
	err = json.Unmarshal(r.ArrJSON, &r.Arr)
	return
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &RefreshArr{})
}
