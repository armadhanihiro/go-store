package repository

import (
	"fmt"
	"gostore/internal/app/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(user model.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		log.Error(fmt.Errorf("error UserRepository - Create : %w", err))
		return err
	}

	return nil
}

func (r *UserRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		log.Error(fmt.Errorf("error UserRepository - FindByemail : %w", err))
		return user, err
	}

	return user, nil
}

func (r *UserRepository) FindByID(id int) (model.User, error) {
	var user model.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		log.Error(fmt.Errorf("error UserRepository - FindByID : %w", err))
		return user, err
	}

	return user, nil
}

func (r *UserRepository) Update(user model.User) error {
	if err := r.db.Updates(user).Error; err != nil {
		log.Error(fmt.Errorf("error UserRepository - Update : %w", err))
		return err
	}

	return nil
}
