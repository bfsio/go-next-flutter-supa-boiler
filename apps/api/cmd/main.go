package main

import (
	"log"
	"net/http"
	"pet-rock-backend/handlers"
	"pet-rock-backend/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin router
	router := gin.Default()

	// Middlewares
	router.Use(middlewares.AuthMiddleware())

	// Routes
	router.POST("/auth/signup", handlers.Signup)
	router.POST("/auth/login", handlers.Login)
	router.POST("/auth/reset-password", handlers.ResetPassword)

	// Protected route: Pet Rock counter
	router.GET("/petrock", handlers.GetPetRock)
	router.POST("/petrock", handlers.UpdatePetRock)

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
