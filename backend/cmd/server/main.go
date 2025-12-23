package main

import (
	"fmt"
	"log"
	"os"

	"github.com/FizzaSadath/flashcard-app-backend/internal/repo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	cfg := repo.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  "disable",
	}

	db, err := repo.NewDatabase(cfg)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")

	sqlDB, _ := db.DB()
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}

	db.AutoMigrate(&repo.CardEntity{})
}
