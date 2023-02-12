package routes

import "github.com/gin-gonic/gin"

func InitRoutes() *gin.Engine {
	r := gin.Default()

	authRoutes(r) // routes/auth.go
	userRoutes(r) // routes/user.go

	return r
}
