package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/net-http/controllers"
	"github.com/net-http/database"
	"github.com/net-http/models"
	// _ "modernc.org/sqlite"
)

func main() {

	database.InitilizeGormDB()
	//auto migrate
	database.Db.AutoMigrate(&models.User{})

	r := gin.Default()
	newUser := controllers.NewUserController()
	authUser := controllers.NewAuthController()

	// newAuth := controllers.NewAuthController()

	r.POST("/signup", authUser.Signup)
	r.GET("/user/alluser", newUser.GetAllUsers)
	r.GET("/user/oneuser/:id", newUser.GetOneUser)
	r.POST("/user/create", newUser.CreateUser)
	r.DELETE("/user/delete/:id", newUser.DeleteUser)
	r.PATCH("/user/update/:id", newUser.UpdateUser)
	// r.POST("/auth/signup", newUser.Signup)

	log.Println("Listening on http://localhost:8080/")

	// http.ListenAndServe(":8080", mux)
	r.Run(":8080")

}
