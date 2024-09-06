package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"not null" json:"name"`
	Age  int    `gorm:"not null" json:"age"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
