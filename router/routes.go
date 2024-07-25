package router

import (
	"github.com/gin-gonic/gin"
	authController "go-to-chat/app/auth"
	"go-to-chat/app/middleware"
	userController "go-to-chat/app/user"
)

func SetupRoutes(router *gin.Engine) {
	apiRouter := router.Group("/api/")
	{
		v1Group := apiRouter.Group("/v1/")
		{
			v1Group.POST("/user", userController.CreateUser)
			v1Group.POST("/auth/login", authController.Login)

			authWithRefreshTokenGroup := v1Group.Group("/")
			authWithRefreshTokenGroup.Use(middleware.AuthHandler(middleware.TokenTypeRefreshToken))
			{
				authWithRefreshTokenGroup.POST("/auth/refresh", authController.RefreshToken)
			}

			authWithAccessTokenGroup := v1Group.Group("/")
			authWithAccessTokenGroup.Use(middleware.AuthHandler(middleware.TokenTypeAccessToken))
			{
				authRouter := authWithAccessTokenGroup.Group("/auth")
				{
					authRouter.POST("/logout", authController.Logout)
				}
				userRouter := authWithAccessTokenGroup.Group("/user")
				{
					userRouter.GET("/:id", userController.GetUser)
					userRouter.PUT("/:id", userController.UpdateUser)
				}
			}
		}
	}
}
