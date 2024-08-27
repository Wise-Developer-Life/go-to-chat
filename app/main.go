package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-to-chat/app/config"
	"go-to-chat/app/job"
	"go-to-chat/app/middleware"
	"go-to-chat/app/socket"
	"go-to-chat/database"
	"go-to-chat/router"
	"log"
)

func main() {
	appConfig, err := config.GetAppConfig()

	if err != nil {
		log.Println("Error getting app config: ", err)
		return
	}

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
	app.GET("/ws", socket.HandleWebSocket())
	app.Use(middleware.ErrorHandler())
	router.SetupRoutes(app)

	err = app.Run(fmt.Sprintf(":%d", appConfig.Server.Port))

	if err != nil {
		return
	}
}
