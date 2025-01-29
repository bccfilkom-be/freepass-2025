package database

import (
	"jevvonn/bcc-be-freepass-2025/internal/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	config := config.GetConfig()

	dsn := config.GetString("connection.database.mysql_uri")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database, %s", err)
	}

	return db
}
