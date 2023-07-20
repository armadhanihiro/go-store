package schema

import (
	"time"
)

type CategoryProduct struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateProductCategory struct {
	Name   	   string         `validate:"required" json:"name"`
}

type DetailProductCategory struct {
	ID          int                 `json:"id"`
	Name        string              `json:"name"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
}

type UpdateProductCategory struct {
	Name   	  string        `validate:"required" json:"name"`
}