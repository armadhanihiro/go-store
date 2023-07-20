package repository

import (
	"fmt"
	"gostore/internal/app/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db}
}

func (r *AuthRepository) Create(auth model.Auth) error {
	if err := r.db.Create(&auth).Error; err != nil {
		log.Error(fmt.Errorf("error AuthRepository - Create : %w", err))
		return err
	}

	return nil
}

func (r *AuthRepository) DeleteAllByUserID(userID int) error {
	auth := model.Auth{}
	if err := r.db.Where("user_id = ?", userID).Delete(&auth).Error; err != nil {
		log.Error(fmt.Errorf("error AuthRepository - DeleteAllByUserID : %w", err))
		return err
	}

	return nil
}

func (ar *AuthRepository) Find(userID int, refreshToken string) (model.Auth, error) {
	var auth model.Auth
	if err := ar.db.Where("user_id = ? AND token = ?", userID, refreshToken).First(&auth).Error; err != nil {
		log.Error(fmt.Errorf("error AuthRepository - Find : %w", err))
		return auth, err
	}

	return auth, nil
}
