package route

import (
	"chapter3-mygram/controller"
	"chapter3-mygram/middleware"
	"chapter3-mygram/repository"
	"chapter3-mygram/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupSocialMediaRoute(router *gin.Engine, db *gorm.DB) {
	socialmediaRepository := repository.NewSocialMediaRepository(db)
	socialmediaService := service.NewSocialService(socialmediaRepository)
	socialmediaController := controller.NewSocialController(socialmediaService)

	authUser := router.Group("/social-media", middleware.AuthMiddleware)
	{
		authUser.POST("", socialmediaController.CreateSocial)
		authUser.GET("", socialmediaController.GetAll)
		authUser.GET("/:social_media_id", socialmediaController.GetOne)
		authUser.PUT("/:social_media_id", socialmediaController.UpdateSocialMedia)
		authUser.DELETE("/:social_media_id", socialmediaController.DeleteSocialMedia)
	}
}
