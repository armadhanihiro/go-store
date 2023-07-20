package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID                int `gorm:"primaryKey;autoIncrement"`
	Name              string
	Description       string
	ProductCategoryID int
	Sold              int
	Amount            int
	Views             int
	Image             string
	CreatedAt         time.Time      `gorm:"autoCreateTime:true;notNull"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime:true;notNull"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
	ProductCategory   ProductCategory
}
