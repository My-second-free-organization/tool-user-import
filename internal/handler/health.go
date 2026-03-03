package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "healthy", "service": "platform-api"}) }
func ReadinessCheck(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ready"}) }
