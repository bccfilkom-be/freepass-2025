package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/litegral/freepass-2025/internal/server"
)

func main() {
	// Load environment variables from file only if not in Docker
	if os.Getenv("DOCKER_ENVIRONMENT") == "" {
		if err := godotenv.Load(); err != nil {
			log.Printf("Warning: Error loading .env file: %v\n", err)
		}
	}

	// Initialize and start server
	srv, err := server.NewServer(os.Getenv("APP_DB_URL"), os.Getenv("APP_URL"))
	if err != nil {
		log.Fatalf("Failed to initialize server: %v\n", err)
	}

	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server: %v\n", err)
	}
}
