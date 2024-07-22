package router

import (
	"github.com/gin-gonic/gin"
	userController "go-to-chat/app/controller"
)

func SetupRoutes(router *gin.Engine) {
	apiRouter := router.Group("/api/")
	{
		v1Router := apiRouter.Group("/v1/")
		{
			userRouter := v1Router.Group("/user/")
			{
				userRouter.GET("/:id", userController.GetUser)
				userRouter.POST("/", userController.CreateUser)
				userRouter.PUT("/:id", userController.UpdateUser)
			}
		}
	}
}
