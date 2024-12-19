package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sohibjon7731/ecommerce_backend/internal/auth/dto"
	"github.com/sohibjon7731/ecommerce_backend/internal/auth/service"
)

type AuthHandler struct {
	Service *service.AuthService
}

func NewAuthHandler() *AuthHandler {
	service := service.NewAuthService()
	return &AuthHandler{Service: service}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var input dto.RegisterDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}
	user, err := h.Service.Register(input.Email, input.Password, input.PasswordConfirmation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{

		"email":      user.Email,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input dto.LoginDTO

	decoder := json.NewDecoder(c.Request.Body)
	decoder.DisallowUnknownFields() // Oshiqcha maydonlarni rad etish

	if err := decoder.Decode(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	accessToken, refreshToken, err := h.Service.Login(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}