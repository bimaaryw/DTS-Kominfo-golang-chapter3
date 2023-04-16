package routers

import (
	"chapter3-sesi2/controllers"
	"chapter3-sesi2/repository"
	"chapter3-sesi2/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupAuthRoute(router *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(*userRepository)
	userController := controllers.NewUserController(*userService)

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", userController.Registration)
		userRouter.POST("/login", userController.Login)
	}
}
