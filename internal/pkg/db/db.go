package db

import (
	"gostore/internal/app/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(dbUri string) (*gorm.DB, error) {
	dsn := dbUri
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}

	db.AutoMigrate(
		&model.User{},
		&model.Auth{},
		&model.ProductCategory{},
		&model.Product{},
		&model.WishList{},
	)

	return db, nil
}
