package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Welcome returns demo API welcome message with requester details
func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Demo API",
		"version": "2.0.0",
		"requester": gin.H{
			"user_agent":    c.GetHeader("User-Agent"),
			"remote_addr":   c.ClientIP(),
			"x_forwarded_for": c.GetHeader("X-Forwarded-For"),
			"x_real_ip":     c.GetHeader("X-Real-IP"),
			"host":          c.GetHeader("Host"),
		},
	})
}

// HealthCheck returns API health status
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "demo-api v2",
	})
}

// GetDocs returns dummy documentation data
func GetDocs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"api":         "Demo API",
		"endpoints":   []string{"/", "/health", "/docs"},
		"description": "Simple demo API for ArgoCD deployment",
	})
}
