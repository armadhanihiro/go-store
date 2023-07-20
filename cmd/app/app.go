package app

import (
	"gostore/internal/app/controller"
	"gostore/internal/app/repository"
	"gostore/internal/app/service"
	"gostore/internal/pkg/config"
	"gostore/internal/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	cfg    config.Config
	DBConn *gorm.DB
	router *gin.Engine
}

func NewServer(cfg config.Config, DBConn *gorm.DB) (*Server, error) {
	server := &Server{
		cfg:    cfg,
		DBConn: DBConn,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	r := gin.New()

	r.Use(
		middleware.LoggingMiddleware(),
		middleware.RecoveryMiddleware(),
	)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ping": "pong",
		})
	})

	// repository
	userRepo := repository.NewUserRepository(server.DBConn)
	authRepo := repository.NewAuthRepository(server.DBConn)
	categoryRepository := repository.NewProductCategoryRepository(server.DBConn)
	productRepository := repository.NewProductRepository(server.DBConn)
	wishListRepository := repository.NewWishListRepository(server.DBConn)
	shoppingChartRepository := repository.NewShoppingChartRepository(server.DBConn)

	// service
	signUpService := service.NewSignUpService(userRepo)
	tokenCreator := service.NewTokenCreator(
		server.cfg.AccessTokenKey,
		server.cfg.RefreshTokenKey,
		server.cfg.AccessTokenDuration,
		server.cfg.RefreshTokenDuration,
	)
	uploadService := service.NewUploaderService(
		server.cfg.CloudinaryCloudName,
		server.cfg.CloudinaryApiKey,
		server.cfg.CloudinaryApiSecret,
		server.cfg.CloudinaryUploadFolder,
	)
	sessionService := service.NewSessionService(userRepo, authRepo, tokenCreator)
	userService := service.NewUserService(userRepo)
	categoryService := service.NewProductCategoryService(categoryRepository)
	productService := service.NewProductService(uploadService, categoryRepository, productRepository)
	wishListService := service.NewWishListService(userRepo, productRepository, wishListRepository)
	shoppingChartService := service.NewShoppingChartService(userRepo, productRepository, shoppingChartRepository)

	// controller
	signUpController := controller.NewSignUpController(signUpService)
	sessionController := controller.NewSessionController(sessionService, tokenCreator)
	userController := controller.NewUserController(userService)
	categoryController := controller.NewCategoryController(categoryService)
	productController := controller.NewProductController(productService)
	wishListController := controller.NewWishListController(wishListService)
	shoppingChartController := controller.NewShoppingChartController(shoppingChartService)

	r.POST("/auth/signup", signUpController.Insert)
	r.POST("/auth/signin", sessionController.SignIn)
	r.GET("/auth/refresh", sessionController.Refresh)
	r.GET("products/", productController.BrowseProduct)
	r.GET("products/:id", productController.DetailProduct)

	r.Use(middleware.AuthMiddleware(tokenCreator))
	r.GET("/auth/signout", sessionController.SignOut)

	r.GET("/users/detail", userController.FindByID)
	r.PATCH("/users", userController.UpdateByID)

	category := r.Group("/categories")
	{
		category.POST("/", categoryController.CreateCategory)
		category.GET("/", categoryController.GetAllCategory)
		category.GET("/:id", categoryController.DetailCategory)
		category.PATCH("/:id", categoryController.UpdateCategory)
		category.DELETE("/:id", categoryController.DeleteCategory)
	}

	products := r.Group("/products")
	{
		products.POST("/", productController.CreateCategory)
		products.PATCH("/:id", productController.UpdateProduct)
		products.DELETE("/:id", productController.DeleteProduct)
	}

	wish_list := r.Group("/wish_list")
	{
		wish_list.POST("/", wishListController.CreateWishList)
		wish_list.GET("/", wishListController.BrowseWishList)
		wish_list.GET("/:id", wishListController.DetailWishList)
		wish_list.PATCH("/:id", wishListController.UpdateWishList)
		wish_list.DELETE("/:id", wishListController.DeleteWishList)
	}

	shopping_chart := r.Group("/shopping_chart")
	{
		shopping_chart.POST("/", shoppingChartController.CreateShoppingChart)
		shopping_chart.GET("/", shoppingChartController.BrowseShoppingChart)
		shopping_chart.GET("/:id", shoppingChartController.DetailShoppingChart)
		shopping_chart.PATCH("/:id", shoppingChartController.UpdateShoppingChart)
		shopping_chart.DELETE("/:id", shoppingChartController.DeleteShoppingChart)
	}

	server.router = r
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
