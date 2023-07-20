package schema

import (
	"time"
)

type CreateWishList struct {
	UserID 		int	`validate:"required" form:"user_id"`
	ProductID 	int	`validate:"required" form:"product_id"`
}

type BrowseWishList struct {
	Product 	string
	User		string
	SortBy   	string
	Page     	int
	Limit    	int
}

type DetailWishList struct {
	ID          int					`json:"id"`
	Views       int                 `json:"views"`
	User    	DetailUserResp 		`json:"user"`
	Product    	DetailProductResp 	`json:"product"`
	CreatedAt   time.Time			`json:"created_at"`
	UpdatedAt   time.Time			`json:"updated_at"`
}

type UpdateWishList struct {
	UserID 			int    `validate:"required" form:"user_id"`
	ProductID 		int    `validate:"required" form:"product_id"`
}
