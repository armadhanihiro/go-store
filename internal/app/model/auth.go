package model

import (
	"time"

	"gorm.io/gorm"
)

type Auth struct {
	gorm.Model
	ID        int `gorm:"primaryKey;autoIncrement"`
	Token     string
	AuthType  string
	ExpiredAt time.Time
	UserID    int
	User      User
}
