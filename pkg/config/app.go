package config

import (
	"fmt"

	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

var db *gorm.DB

func ConnectDB() {

	database, err := gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	fmt.Println("Database Connection Successful ✅")

	db = database
}

func GetDB() *gorm.DB {
	return db
}
