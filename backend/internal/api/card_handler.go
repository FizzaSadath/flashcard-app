package api

import (
	"net/http"
	"strconv"

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

func (h *CardHandler) GetStats(c *gin.Context) {
	userID := getUserID(c)
	stats, err := h.service.GetStats(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch stats"})
		return
	}
	c.JSON(http.StatusOK, stats)
}

func (h *CardHandler) GetDeckStats(c *gin.Context) {
	userID := getUserID(c)
	stats, err := h.service.GetDeckStats(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stats)
}

func (h *CardHandler) DeleteCard(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 32)
	cardID := uint(id)

	userID := getUserID(c)

	err := h.service.DeleteCard(userID, cardID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Card deleted"})
}

func (h *CardHandler) UploadCSV(c *gin.Context) {
	deckIDStr := c.Param("id")
	id64, err := strconv.ParseUint(deckIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid deck ID"})
		return
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer file.Close()

	userID := getUserID(c)

	count, err := h.service.ImportCards(userID, uint(id64), file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Import successful", "count": count})
}
