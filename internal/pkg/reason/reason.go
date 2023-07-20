package reason

var (
	InternalServerError = "internal server error"
	RequestFormError    = "request format is not valid"
)

var (
	UserAlreadyExist   = "user already exist"
	SignUpFailed       = "failed to sign up"
	Unauthorized       = "unauthorized request"
	UserNotFound       = "user not exist"
	FailedSignIn       = "failed to sign in, your email or password is incorrect"
	FailedSignOut      = "failed to sign out"
	FailedRefreshToken = "failed to refresh token, please check your token" //nolint
	FailedUpdateUser   = "failed to update user"
)

var (
	CategoryNotFound        = "category not found"
	CategoryFailedCreate    = "failed to create category"
	CategoryFailedBrowse    = "failed to get categories"
	CategoryFailedUpdate    = "failed to update category"
	CategoryFailedDelete    = "failed to delete category"
	CategoryFailedGetDetail = "failed to get category detail"
)

var (
	ProductNotFound          = "product not found"
	ProductFailedCreate      = "failed to create product"
	ProductFailedBrowse      = "failed to get products"
	ProductFailedUpdate      = "failed to update product"
	ProductFailedDelete      = "failed to delete product"
	ProductImageFailedUpload = "failed to upload product image"
)

var (
	WishListNotFound     = "wishList not found"
	WishListFailedCreate = "failed to create wishList"
	WishListFailedBrowse = "failed to get wishLists"
	WishListFailedUpdate = "failed to update wishList"
	WishListFailedDelete = "failed to delete wishList"
)

var (
	ShoppingChartNotFound     = "shopping cart not found"
	ShoppingChartFailedCreate = "failed to create shopping cart"
	ShoppingChartFailedBrowse = "failed to get shopping cart"
	ShoppingChartFailedUpdate = "failed to update shopping cart"
	ShoppingChartFailedDelete = "failed to delete shopping cart"
)
