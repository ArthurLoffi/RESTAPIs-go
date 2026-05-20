package user

import (
	"time"
)

type User struct {
	ID uint `gorm:"primaryKey;autoincrement"`
	Name string `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

var Users = []User{}