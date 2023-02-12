package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/net-http/controllers"
)

func userRoutes(e *gin.Engine) {
	userApi := controllers.NewUserController()

	userGroup := e.Group("/user")
	{
		userGroup.GET("", userApi.GetAllUsers)
		userGroup.POST("", userApi.CreateUser)
		userGroup.GET("/:id", userApi.GetOneUser)
		userGroup.DELETE("/:id", userApi.DeleteUser)
		userGroup.PATCH("/:id", userApi.UpdateUser)
	}
}
