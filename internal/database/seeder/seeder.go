package main

import (
	"fmt"
	"jevvonn/bcc-be-freepass-2025/internal/constant"
	"jevvonn/bcc-be-freepass-2025/internal/database"
	"jevvonn/bcc-be-freepass-2025/internal/models/domain"
)

func main() {
	db := database.NewDatabase()

	fmt.Println("Setuping data seeder...")
	fmt.Println("Seeding data...")
	db.Create(&domain.User{
		Name:     "Admin",
		Email:    "admin@gmail.com",
		Password: "$2a$12$4kNcMNzbImfqZvcCf.LEoOl2meaVjs2pTpb8SrXh2BJ29A6bWV95W", // password
		Role:     constant.ROLE_ADMIN,
		Bio:      "Aku Admin",
	})

	db.Create(&domain.User{
		Name:     "Kordinator Event",
		Email:    "kordinator@gmail.com",
		Password: "$2a$12$4kNcMNzbImfqZvcCf.LEoOl2meaVjs2pTpb8SrXh2BJ29A6bWV95W", // password
		Role:     constant.ROLE_COORDINATOR,
		Bio:      "Aku Kordinator Event",
	})

	db.Create(&domain.User{
		Name:     "User1",
		Email:    "user1@gmail.com",
		Password: "$2a$12$4kNcMNzbImfqZvcCf.LEoOl2meaVjs2pTpb8SrXh2BJ29A6bWV95W", // password
		Role:     constant.ROLE_USER,
	})
	fmt.Println("Data seeded")

}
