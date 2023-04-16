package main

import (
	"chapter3-sesi2/database"
	"chapter3-sesi2/routers"

	"github.com/gin-gonic/gin"
)

const PORT = ":8080"

func main() {
	router := gin.Default()

	database.StartDB()
	db := database.GetDB()

	routers.SetupAuthRoute(router, db)
	routers.SetupProductRoute(router, db)

	router.Run(PORT)
}
