package main

import (
	"github.com/gin-gonic/gin"
	"go-to-chat/app/middleware"
	"go-to-chat/database"
	"go-to-chat/router"
)

func main() {
	database.SetupDatabase()

	app := gin.Default()
	app.Use(gin.Logger())
	app.Use(middleware.ErrorHandler())

	router.SetupRoutes(app)

	err := app.Run(":8080")
	if err != nil {
		return
	}
}
