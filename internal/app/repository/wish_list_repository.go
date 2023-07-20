package repository

import (
	"fmt"
	"gostore/internal/app/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type WishListRepository struct {
	db *gorm.DB
}

func NewWishListRepository(db *gorm.DB) *WishListRepository {
	return &WishListRepository{db}
}

func (r *WishListRepository) Create(wish_list model.WishList) error {
	if err := r.db.Create(&wish_list).Error; err != nil {
		log.Error(fmt.Errorf("error WishListRepository - Create : %w", err))
		return err
	}
	return nil
}

func (r *WishListRepository) FindAll(sort_by string, page int, limit int) ([]model.WishList, error) {
	wish_list := []model.WishList{}
	offset := limit * (page - 1)
	sort := fmt.Sprintf("wish_list.%s DESC", sort_by)

	err := r.db.Joins("left join user ON user.id = user_id").Joins("left join product ON product.id = product_id").Order(sort).Order(sort).Limit(limit).Offset(offset).Find(&wish_list).Error
	if err != nil {
		return wish_list, err
	}

	return wish_list, nil
}

func (r *WishListRepository) FindByID(id int) (model.WishList, error) {
	wish_list := model.WishList{}
	if err := r.db.Joins("User").Joins("Product").Where("wish_list.id = ?", id).First(&wish_list).Error; err != nil {
		log.Error(fmt.Errorf("error WishListRepository - FindByID : %w", err))
		return wish_list, err
	}

	return wish_list, nil
}

func (r *WishListRepository) Update(wish_list model.WishList) error {
	if err := r.db.Updates(&wish_list).Error; err != nil {
		log.Error(fmt.Errorf("error WishListRepository - Update : %w", err))
		return err
	}

	return nil
}

func (r *WishListRepository) DeleteByID(id int) error {
	wish_list := model.WishList{}
	if err := r.db.Where("id = ?", id).Delete(&wish_list).Error; err != nil {
		log.Error(fmt.Errorf("error WishListRepository - DeleteByID : %w", err))
		return err
	}
	return nil
}
