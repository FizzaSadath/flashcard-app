package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(cardHandler *CardHandler, authHandler *AuthHandler, deckHandler *DeckHandler) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Nuxt URL
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// Health Check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	api := r.Group("/api")
	{
		// Auth Routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// Card Routes
		protected := api.Group("/")
		protected.Use(AuthMiddleware())
		{
			protected.POST("/decks", deckHandler.CreateDeck)
			protected.GET("/decks", deckHandler.ListDecks)
			protected.DELETE("/decks/:id", deckHandler.DeleteDeck)

			protected.POST("/cards", cardHandler.CreateCard)
			protected.GET("/cards", cardHandler.ListCards)
			protected.GET("/cards/due", cardHandler.ListDueCards)
			protected.POST("/cards/review", cardHandler.ReviewCard)
			protected.GET("/stats", cardHandler.GetStats)
		}
	}

	return r
}
