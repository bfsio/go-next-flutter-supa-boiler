package handlers

import (
	"fmt"
	"net/http"
	"pet-rock-backend/services"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var req services.SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	user, err := services.SignupUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Signup failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func Login(c *gin.Context) {
	var req services.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	token, err := services.LoginUser(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func ResetPassword(c *gin.Context) {
	var req services.ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	err := services.ResetPassword(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password reset failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset link sent"})
}
