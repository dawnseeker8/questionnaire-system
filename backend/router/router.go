package router

import (
	"time"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cors"
)

// RegisterRoutes registers all the routes for the application
func RegisterRoutes(h *server.Hertz) {
	// Configure CORS
	h.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// API v1 group
	v1 := h.Group("/api/v1")
	
	// Health check endpoint
	v1.GET("/health", HealthCheck)

	// Register other route groups
	setupQuestionnaireRoutes(v1)
}
