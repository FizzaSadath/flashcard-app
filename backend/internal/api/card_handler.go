package api

import (
	"net/http"

	"github.com/FizzaSadath/flashcard-app-backend/internal/core"
	"github.com/gin-gonic/gin"
)

type CardHandler struct {
	service *core.CardService
}

func NewCardHandler(service *core.CardService) *CardHandler {
	return &CardHandler{service: service}
}

func getUserID(c *gin.Context) uint {
	return c.MustGet("userID").(uint)
}

func (h *CardHandler) CreateCard(c *gin.Context) {
	var req struct {
		Front  string `json:"front" binding:"required"`
		Back   string `json:"back" binding:"required"`
		DeckID uint   `json:"deck_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := getUserID(c)

	card, err := h.service.CreateCard(userID, req.DeckID, req.Front, req.Back)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, card)
}

func (h *CardHandler) ListCards(c *gin.Context) {
	userID := getUserID(c)

	cards, err := h.service.ListCards(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cards"})
		return
	}

	c.JSON(http.StatusOK, cards)
}

func (h *CardHandler) ReviewCard(c *gin.Context) {
	var req struct {
		CardID uint `json:"card_id" binding:"required"`
		Grade  int  `json:"grade" binding:"required,min=0,max=5"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := getUserID(c)

	err := h.service.ReviewCard(userID, req.CardID, req.Grade)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review recorded successfully"})
}

func (h *CardHandler) ListDueCards(c *gin.Context) {
	userID := getUserID(c)

	cards, err := h.service.ListDueCards(userID, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch due cards"})
		return
	}

	c.JSON(http.StatusOK, cards)
}
