package services

import (
	"github.com/gin-gonic/gin"
)

func GetHealth(c *gin.Context) {
	c.JSON(201, gin.H{
		"status": "healthy",
	})
}
