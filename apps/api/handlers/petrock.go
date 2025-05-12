package handlers

import (
	"fmt"
	"net/http"
	"pet-rock-backend/services"
	"github.com/gin-gonic/gin"
)

func GetPetRock(c *gin.Context) {
	tenantUUID := c.Param("tenant_uuid") // Extract tenant UUID
	userUUID := c.MustGet("user_uuid").(string) // User ID from authenticated JWT

	counter, err := services.GetPetRockCounter(userUUID, tenantUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch pet rock counter"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"counter": counter})
}

func UpdatePetRock(c *gin.Context) {
	tenantUUID := c.Param("tenant_uuid") // Extract tenant UUID
	userUUID := c.MustGet("user_uuid").(string) // User ID from authenticated JWT

	var req services.UpdatePetRockRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	err := services.UpdatePetRockCounter(userUUID, tenantUUID, req.Count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update pet rock counter"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Counter updated successfully"})
}
