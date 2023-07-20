package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             int `gorm:"primaryKey;autoIncrement"`
	FullName       string
	Email          string
	HashedPassword string
	Address        string
	City           string
	Province       string
	Country        string
	PostalCode     string
	CreatedAt      time.Time      `gorm:"autoCreateTime:true;notNull"`
	UpdatedAt      time.Time      `gorm:"autoUpdateTime:true;notNull"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
