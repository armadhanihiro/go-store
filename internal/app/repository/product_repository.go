package repository

import (
	"fmt"
	"gostore/internal/app/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (r *ProductRepository) Create(product model.Product) error {
	if err := r.db.Create(&product).Error; err != nil {
		log.Error(fmt.Errorf("error ProductRepository - Create : %w", err))
		return err
	}

	return nil
}

func (r *ProductRepository) FindAll(name string, sort_by string, page int, limit int) ([]model.Product, error) {
	products := []model.Product{}
	offset := limit * (page - 1)
	nameLike := fmt.Sprintf("%s%s%s", "%", name, "%")
	sort := fmt.Sprintf("products.%s DESC", sort_by)

	err := r.db.Joins("left join product_categories ON product_categories.id = product_category_id").Preload("ProductCategory").Where("products.name ILIKE ?", nameLike).Order(sort).Order(sort).Limit(limit).Offset(offset).Find(&products).Error
	if err != nil {
		return products, err
	}

	return products, nil
}

func (r *ProductRepository) FindByID(id int) (model.Product, error) {
	product := model.Product{}
	if err := r.db.Joins("ProductCategory").Where("products.id = ?", id).First(&product).Error; err != nil {
		log.Error(fmt.Errorf("error ProductRepository - FindByID : %w", err))
		return product, err
	}

	return product, nil
}

func (r *ProductRepository) Update(product model.Product) error {
	if err := r.db.Updates(&product).Error; err != nil {
		log.Error(fmt.Errorf("error ProductRepository - Update : %w", err))
		return err
	}

	return nil
}

func (r *ProductRepository) UpdateImageURL(id int, imageURL string) error {
	return nil
}

func (r *ProductRepository) DeleteByID(id int) error {
	product := model.Product{}
	if err := r.db.Where("id = ?", id).Delete(&product).Error; err != nil {
		log.Error(fmt.Errorf("error ProductRepository - DeleteByID : %w", err))
		return err
	}
	return nil
}
