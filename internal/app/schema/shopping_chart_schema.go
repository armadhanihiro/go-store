package schema

import (
	"time"
)

type CreateShoppingChart struct {
	UserID 		int	`validate:"required" form:"user_id"`
	ProductID 	int	`validate:"required" form:"product_id"`
	Quantity	int`validate:"required" form:"quantity"`
}

type BrowseShoppingChart struct {
	Product 	string
	User		string
}

type DetailShoppingChart struct {
	ID          int					`json:"id"`
	Quantity    int                 `json:"quantity"`
	User    	DetailUserResp 		`json:"user"`
	Product    	DetailProductResp 	`json:"product"`
	CreatedAt   time.Time			`json:"created_at"`
	UpdatedAt   time.Time			`json:"updated_at"`
}

type UpdateShoppingChart struct {
	UserID 		int	`validate:"required" form:"user_id"`
	ProductID 	int	`validate:"required" form:"product_id"`
	Quantity	int`validate:"required" form:"quantity"`
}
