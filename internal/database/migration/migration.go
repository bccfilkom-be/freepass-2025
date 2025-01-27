package main

import (
	"fmt"
	"jevvonn/bcc-be-freepass-2025/internal/database"
	"jevvonn/bcc-be-freepass-2025/internal/models/domain"
)

func main() {
	db := database.NewDatabase()
	tables := []interface{}{
		&domain.Session{},
		&domain.User{},
	}

	fmt.Println("Setting Up New Migration...")
	fmt.Println("Dropping Tables...")
	db.Migrator().DropTable(tables...)

	fmt.Println("Migrating Tables...")
	db.AutoMigrate(tables...)

	fmt.Println("Migration Completed!")
}
