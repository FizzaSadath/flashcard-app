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
		godotenv.Load("../.env")
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
	db.AutoMigrate(&repo.UserEntity{}, &repo.DeckEntity{}, &repo.CardEntity{})

	cardRepo := repo.NewCardRepo(db)
	userRepo := repo.NewUserRepo(db)
	deckRepo := repo.NewDeckRepo(db)

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "super-secret-key-87$4%1@*7"
	}
	authService := core.NewAuthService(userRepo, jwtSecret)
	deckService := core.NewDeckService(deckRepo)
	cardService := core.NewCardService(cardRepo, deckRepo)

	authHandler := api.NewAuthHandler(authService)
	deckHandler := api.NewDeckHandler(deckService)
	cardHandler := api.NewCardHandler(cardService)

	router := api.SetupRouter(cardHandler, authHandler, deckHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
