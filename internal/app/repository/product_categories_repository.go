package repository

import (
	"fmt"
	"gostore/internal/app/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductCategory struct {
	db *gorm.DB
}

func NewProductCategoryRepository(db *gorm.DB) *ProductCategory {
	return &ProductCategory{db}
}

func (r *ProductCategory) Create(productCategory model.ProductCategory) error {
	if err := r.db.Create(&productCategory).Error; err != nil {
		log.Error(fmt.Errorf("error CategoryRepository - Create : %w", err))
		return err
	}

	return nil
}

func (r *ProductCategory) FindAll() ([]model.ProductCategory, error) {
	category := []model.ProductCategory{}

	if err := r.db.Find(&category).Error; err != nil {
		log.Error(fmt.Errorf("error CategoryRepository - FindAll : %w", err))
		return category, err
	}

	return category, nil
}

func (r *ProductCategory) FindByID(id int) (model.ProductCategory, error) {
	var category = model.ProductCategory{}

	if err := r.db.Where("id = ?", id).First(&category).Error; err != nil {
		log.Error(fmt.Errorf("error CategoryRepository - FindByID : %w", err))
		return category, err
	}

	return category, nil
}

func (r *ProductCategory) FindByName(name string) (model.ProductCategory, error) {
	var category = model.ProductCategory{}

	if err := r.db.Where("name = ?", name).First(&category).Error; err != nil {
		log.Error(fmt.Errorf("error CategoryRepository - FindByName : %w", err))
		return category, err
	}

	return category, nil
}

func (r *ProductCategory) Update(category model.ProductCategory) error {
	if err := r.db.Updates(category).Error; err != nil {
		log.Error(fmt.Errorf("error CategoryRepository - Update : %w", err))
		return err
	}
	return nil
}
func (r *ProductCategory) DeleteByID(id int) error {
	if err := r.db.Where("id = ?", id).Delete(&ProductCategory{}).Error; err != nil {
		log.Error(fmt.Errorf("error CategoryRepository - Delete : %w", err))
		return err
	}

	return nil
}
