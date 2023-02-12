package main

import (
	"log"

	"github.com/net-http/database"
	"github.com/net-http/models"
	"github.com/net-http/routes"
	// _ "modernc.org/sqlite"
)

func main() {

	database.InitilizeGormDB()
	//auto migrate
	database.Db.AutoMigrate(&models.User{})

	r := routes.InitRoutes()

	log.Println("Listening on http://localhost:8080/")
	// http.ListenAndServe(":8080", mux)
	r.Run(":8080")

}
