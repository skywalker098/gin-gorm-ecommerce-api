package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/net-http/controllers"
)

func authRoutes(e *gin.Engine) {
	authApi := controllers.NewAuthController()

	auth := e.Group("/auth")
	{
		auth.POST("/signup", authApi.Signup)
		auth.POST("/login", authApi.Login)
		auth.PATCH("/verify/:email", authApi.VerifyUser)
	}
}
