package main

import (
	"github.com/gin-gonic/gin"
	"go-to-chat/app/chat"
	"go-to-chat/app/job"
	"go-to-chat/app/middleware"
	"go-to-chat/database"
	"go-to-chat/router"
	"log"
)

func main() {
	database.SetupDatabase()

	client := job.GetClientInstance()
	defer func() {
		if err := client.Close(); err != nil {
			log.Fatalf("failed to close client: %v", err)
		}
	}()

	jobRunner := job.CreateJobServer()
	defer jobRunner.Shutdown()

	app := gin.Default()

	app.GET("/ws", chat.HandleChatSocket())
	app.Use(middleware.ErrorHandler())
	router.SetupRoutes(app)

	err := app.Run(":8082")
	if err != nil {
		return
	}
}
