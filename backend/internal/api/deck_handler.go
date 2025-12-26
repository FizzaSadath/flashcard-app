package api

import (
	"strconv"

	"github.com/FizzaSadath/flashcard-app-backend/internal/core"
	"github.com/gin-gonic/gin"
)

type DeckHandler struct {
	service *core.DeckService
}

func NewDeckHandler(service *core.DeckService) *DeckHandler {
	return &DeckHandler{service: service}
}

func (h *DeckHandler) CreateDeck(c *gin.Context) {
	var req struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	userID := getUserID(c)

	deck, err := h.service.CreateDeck(userID, req.Name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, deck)
}

func (h *DeckHandler) ListDecks(c *gin.Context) {
	userID := getUserID(c)
	decks, err := h.service.ListDecks(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, decks)
}

func (h *DeckHandler) DeleteDeck(c *gin.Context) {
	userID := getUserID(c)
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	err := h.service.DeleteDeck(userID, uint(id))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Deck deleted"})
}
