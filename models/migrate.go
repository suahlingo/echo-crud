package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID    int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"not null" json:"name"`
	Email string `gorm:"not null;unique" json:"email"`
	Age   int    `gorm:"not null" json:"age"`
}

func Migrate(db *gorm.DB) error {
}
