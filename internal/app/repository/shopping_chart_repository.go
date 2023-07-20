package repository

import (
	"fmt"
	"gostore/internal/app/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ShoppingChartRepository struct {
	db *gorm.DB
}

func NewShoppingChartRepository(db *gorm.DB) *ShoppingChartRepository {
	return &ShoppingChartRepository{db}
}

func (r *ShoppingChartRepository) Create(shopping_chart model.ShoppingChart) error {
	if err := r.db.Create(&shopping_chart).Error; err != nil {
		log.Error(fmt.Errorf("error ShoppingChartRepository - Create : %w", err))
		return err
	}
	return nil
}

func (r *ShoppingChartRepository) FindAll() ([]model.ShoppingChart, error) {
	shopping_chart := []model.ShoppingChart{}

	err := r.db.Joins("left join user ON user.id = user_id").Joins("left join product ON product.id = product_id").Find(&shopping_chart).Error
	if err != nil {
		return shopping_chart, err
	}

	return shopping_chart, nil
}

func (r *ShoppingChartRepository) FindByID(id int) (model.ShoppingChart, error) {
	shopping_chart := model.ShoppingChart{}
	if err := r.db.Joins("User").Joins("Product").Where("shopping_chart.id = ?", id).First(&shopping_chart).Error; err != nil {
		log.Error(fmt.Errorf("error ShoppingChartRepository - FindByID : %w", err))
		return shopping_chart, err
	}

	return shopping_chart, nil
}

func (r *ShoppingChartRepository) Update(shopping_chart model.ShoppingChart) error {
	if err := r.db.Updates(&shopping_chart).Error; err != nil {
		log.Error(fmt.Errorf("error ShoppingChartRepository - Update : %w", err))
		return err
	}

	return nil
}

func (r *ShoppingChartRepository) DeleteByID(id int) error {
	shopping_chart := model.ShoppingChart{}
	if err := r.db.Where("id = ?", id).Delete(&shopping_chart).Error; err != nil {
		log.Error(fmt.Errorf("error ShoppingChartRepository - DeleteByID : %w", err))
		return err
	}
	return nil
}
