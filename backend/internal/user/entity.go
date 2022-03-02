package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID              uint32 `gorm:"primaryKey"`
	Username        string
	Password        string
	Nickname        string
	SelfDescription string
	Avatar          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}
