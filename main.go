package main

import (
	"github.com/gin-gonic/gin"
	"go-to-chat/app/chat"
	"go-to-chat/app/middleware"
	"go-to-chat/database"
	"go-to-chat/router"
)

func main() {
	database.SetupDatabase()

	app := gin.Default()

	// FIXME: This is a security risk, do not use this in production
	app.StaticFile("/", "./home.html")
	app.GET("/ws", chat.HandleChatSocket())

	app.Use(gin.Logger())
	app.Use(middleware.ErrorHandler())

	router.SetupRoutes(app)

	err := app.Run(":8082")
	if err != nil {
		return
	}
}
