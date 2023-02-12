package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

var Db *gorm.DB

func InitilizeGormDB() {
	var err error
	// Db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	// if err != nil {
	// 	panic(err)
	// }
	dsn := "host=localhost user=postgres password=rohan123 dbname=postgres port=3030 sslmode=disable"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Unable to connect to database")

	}
}
