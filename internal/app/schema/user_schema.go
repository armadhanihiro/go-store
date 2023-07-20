package schema

import (
	"time"

	"gorm.io/gorm"
)

type SignUpReq struct {
	FullName string `json:"full_name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,alphanum"`
}

type DetailUserResp struct {
	ID         int            `json:"id"`
	FullName   string         `json:"full_name"`
	Email      string         `json:"email"`
	Address    string         `json:"address"`
	City       string         `json:"city"`
	Province   string         `json:"province"`
	Country    string         `json:"country"`
	PostalCode string         `json:"postal_code"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

type UpdateUserReq struct {
	FullName   string `json:"full_name" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,min=8,alphanum"`
	Address    string `json:"address" validate:"required"`
	City       string `json:"city" validate:"required"`
	Province   string `json:"province" validate:"required"`
	Country    string `json:"country" validate:"required"`
	PostalCode string `json:"postal_code" validate:"required"`
}
