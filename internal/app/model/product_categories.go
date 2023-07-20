package model

import (
	"time"

	"gorm.io/gorm"
)

type ProductCategory struct {
	ID        	int				`gorm:"primaryKey;autoIncrement"`
	Name  		string
	CreatedAt 	time.Time		`gorm:"autoCreateTime:true;notNull"`
	UpdatedAt 	time.Time		`gorm:"autoUpdateTime:true;notNull"`
	DeletedAt 	gorm.DeletedAt	`gorm:"index"`
}