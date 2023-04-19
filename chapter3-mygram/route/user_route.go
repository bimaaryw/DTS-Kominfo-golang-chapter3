package route

import (
	"chapter3-mygram/controller"
	"chapter3-mygram/middleware"
	"chapter3-mygram/repository"
	"chapter3-mygram/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupUserRoute(router *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	photoRepository := repository.NewPhotoRepository(db)
	socialRepository := repository.NewSocialMediaRepository(db)
	userService := service.NewUserService(userRepository, photoRepository, socialRepository)
	userController := controller.NewUserController(userService)

	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)

	authUser := router.Group("/user", middleware.AuthMiddleware)
	{
		authUser.GET("/profile", userController.GetProfile)
	}
}
