package main

import (
	"log"

	_ "chapter3-mygram/docs"

	"chapter3-mygram/database"
	"chapter3-mygram/route"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const PORT = ":8080"

// @title					Hacktiv8 Mygram API
// @version					1.0
// @description				Mygram API untuk last project dari DTS dan Hacktiv8.
// @host 					localhost:8080
// @BasePath 				/
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	router := gin.Default()

	database.StartDB()
	db := database.GetDB()

	route.SetupUserRoute(router, db)
	route.SetupPhotoRoute(router, db)
	route.SetupSocialMediaRoute(router, db)
	route.SetupCommentRoute(router, db)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(PORT)
}
