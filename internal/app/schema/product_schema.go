package schema

import (
	"mime/multipart"
	"time"
)

type CreateProductReq struct {
	Name              string                `validate:"required" form:"name"`
	Description       string                `validate:"required" form:"description"`
	ProductCategoryID int                   `validate:"required" form:"product_category_id"`
	Amount            int                   `validate:"required" form:"amount"`
	Image             *multipart.FileHeader `validate:"required,omitempty" form:"image"`
}

type BrowsProductReq struct {
	Name     string
	Category string
	SortBy   string
	Page     int
	Limit    int
}

type DetailProductResp struct {
	ID          int                   `json:"id"`
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Sold        int                   `json:"sold"`
	Amount      int                   `json:"amount"`
	Views       int                   `json:"views"`
	Image       string                `json:"image"`
	Category    DetailProductCategory `json:"category"`
	CreatedAt   time.Time             `json:"created_at"`
	UpdatedAt   time.Time             `json:"updated_at"`
}

type UpdateProductReq struct {
	Name              string `validate:"required" form:"name"`
	Description       string `validate:"required" form:"description"`
	ProductCategoryID int    `validate:"required" form:"product_category_id"`
	Sold              int
	Views             int
	Amount            int                   `validate:"required" form:"amount"`
	Image             *multipart.FileHeader `validate:"required,omitempty" form:"image"`
}
