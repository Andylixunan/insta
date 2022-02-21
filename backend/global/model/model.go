package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint32 `gorm:"primaryKey"`
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
