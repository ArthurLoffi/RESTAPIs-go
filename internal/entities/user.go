package user

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	ID uint `gorm:"primaryKey;autoincrement"`
	Name string `gorm:"size:255;not null"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

var Users = []User{}