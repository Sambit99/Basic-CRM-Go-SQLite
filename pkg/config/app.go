package config

import (
	"fmt"

	"github.com/Sambit99/Basic-CRM-Go-SQLite/pkg/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {

	database, err := gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Database Connection Successful ✅")

	database.AutoMigrate(&model.Lead{})
	fmt.Println("Database Migrated")

	db = database
}

func GetDB() *gorm.DB {
	return db
}
