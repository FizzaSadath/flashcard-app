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
	return &CardHandler{
		service: service,
	}
}

type ReviewRequest struct {
	CardID uint `json:"card_id" binding:"required"`
	Grade  int  `json:"grade" binding:"required,min=0,max=5"`
}

func (h *CardHandler) ReviewCard(c *gin.Context) {
	var req ReviewRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.ReviewCard(req.CardID, req.Grade)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review recorded successfully"})
}

type CreateCardRequest struct {
	Front string `json:"front"`
	Back  string `json:"back"`
}

func (h *CardHandler) CreateCard(c *gin.Context) {
	var req CreateCardRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCard, err := h.service.CreateCard(req.Front, req.Back)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create card"})
		return
	}

	c.JSON(http.StatusCreated, createdCard)
}

func (h *CardHandler) ListCards(c *gin.Context) {
	cards, err := h.service.ListCards()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch cards"})
		return
	}
	c.JSON(http.StatusOK, cards)
}
