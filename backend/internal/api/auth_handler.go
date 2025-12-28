package api

import (
	"net/http"

	"github.com/FizzaSadath/flashcard-app-backend/internal/core"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *core.AuthService
}

func NewAuthHandler(service *core.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

type AuthRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
type RegRequest struct {
	Email           string `json:"email" binding:"required,email"`
	Username        string `json:"username" binding:"required,min=3"`
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,min=6"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Register(req.Email, req.Username, req.Password, req.ConfirmPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req AuthRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":   token,
		"message": "Login successful",
	})
}
