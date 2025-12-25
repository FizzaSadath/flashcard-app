package main

import (
	"log"
	"os"

	"github.com/FizzaSadath/flashcard-app-backend/internal/api"
	"github.com/FizzaSadath/flashcard-app-backend/internal/core"
	"github.com/FizzaSadath/flashcard-app-backend/internal/repo"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system vars")
	}

	dbConfig := repo.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  "disable",
	}

	db, err := repo.NewDatabase(dbConfig)
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	db.AutoMigrate(&repo.UserEntity{}, &repo.CardEntity{})

	cardRepo := repo.NewCardRepo(db)
	cardService := core.NewCardService(cardRepo)
	cardHandler := api.NewCardHandler(cardService)

	router := api.SetupRouter(cardHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
