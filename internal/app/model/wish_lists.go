package model

import (
	"time"

	"gorm.io/gorm"
)

type WishList struct {
	gorm.Model
	ID				int `gorm:"primaryKey;autoIncrement"`
	UserID			int
	ProductID		int
	Views			int
	CreatedAt		time.Time      `gorm:"autoCreateTime:true;notNull"`
	UpdatedAt		time.Time      `gorm:"autoUpdateTime:true;notNull"`
	DeletedAt		gorm.DeletedAt `gorm:"index"`
	User			User
	Product   		Product
}