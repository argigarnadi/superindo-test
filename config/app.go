package config

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"superindo-test/controller"
	"superindo-test/repositories"
	"superindo-test/routes"
	"superindo-test/service"
)

type AppConfig struct {
	DB         *gorm.DB
	Server     *gin.Engine
	LoadConfig *viper.Viper
}

func App(config *AppConfig) {
	// repository layer
	userRepository := repositories.NewUserRepository(config.DB)
	productCategoryRepository := repositories.NewProductCategoryRepository(config.DB)
	productRepository := repositories.NewProductRepository(config.DB, productCategoryRepository)
	cartRepository := repositories.NewCartRepository(config.DB)

	// service layer
	userService := service.NewUserService(userRepository, config.LoadConfig)
	productCategoryService := service.NewProductCategoryService(productCategoryRepository)
	productService := service.NewProductService(productRepository, productCategoryRepository)
	cartService := service.NewCartService(cartRepository, productRepository)

	// controller layer
	userController := controller.NewUserController(userService)
	productCategoryController := controller.NewProductCategoryController(productCategoryService)
	productController := controller.NewProductController(productService)
	cartController := controller.NewCartController(cartService)

	// route config
	routeConfig := routes.RouterConfig{
		Server:                    config.Server,
		LoadConfig:                config.LoadConfig,
		UserController:            userController,
		ProductCategoryController: productCategoryController,
		ProductController:         productController,
		CartController:            cartController,
	}

	routeConfig.Route()
}
