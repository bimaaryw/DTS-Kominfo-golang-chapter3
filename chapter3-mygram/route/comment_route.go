package route

import (
	"chapter3-mygram/controller"
	"chapter3-mygram/middleware"
	"chapter3-mygram/repository"
	"chapter3-mygram/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupCommentRoute(router *gin.Engine, db *gorm.DB) {
	commentRepository := repository.NewCommentRepository(db)
	photoRepository := repository.NewPhotoRepository(db)
	commentService := service.NewCommentService(commentRepository, photoRepository)
	commentController := controller.NewCommentController(commentService)

	authUser := router.Group("/comments", middleware.AuthMiddleware)
	{
		authUser.POST("/:photo_id", commentController.CreateComment)
		authUser.GET("", commentController.GetAll)
		authUser.GET("/:comment_id", commentController.GetOne)
		authUser.PUT("/:comment_id", commentController.UpdateComment)
		authUser.DELETE("/:comment_id", commentController.DeleteComment)
	}
}
