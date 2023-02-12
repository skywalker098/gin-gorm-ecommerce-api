package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/net-http/controllers"
	"github.com/net-http/middleware"
)

func userRoutes(e *gin.Engine) {
	e.Use(middleware.AuthMiddleware())
	userApi := controllers.NewUserController()
	userGroup := e.Group("/user")
	{
		//to use this endpoint, provide the bearer token from login response
		//in the authorisation header - bearer token

		userGroup.GET("", userApi.GetAllUsers)
		userGroup.POST("", userApi.CreateUser)
		userGroup.GET("/:id", userApi.GetOneUser)
		userGroup.DELETE("/:id", userApi.DeleteUser)
		userGroup.PATCH("/:id", userApi.UpdateUser)
	}
}
