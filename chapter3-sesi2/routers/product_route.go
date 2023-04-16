package routers

import (
	"chapter3-sesi2/controllers"
	"chapter3-sesi2/middleware"
	"chapter3-sesi2/repository"
	"chapter3-sesi2/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupProductRoute(router *gin.Engine, db *gorm.DB) {
	productRepository := repository.NewProductRepository(db)
	userRepository := repository.NewUserRepository(db)
	productService := services.NewProductService(*productRepository, *userRepository)
	productController := controllers.NewProductController(*productService)

	productRouter := router.Group("/product", middleware.AuthMiddleware)
	{
		productRouter.POST("/", productController.CreateProduct)
		productRouter.GET("/", productController.GetProduct)
		adminRouter := productRouter.Group("/", middleware.AdminMiddleware)
		{
			adminRouter.PUT(":product_id", productController.UpdateProduct)
			adminRouter.DELETE(":product_id", productController.DeleteProduct)
		}
	}
}

func test(c *gin.Context) {
	c.JSON(http.StatusOK, "Berhasil")
}
