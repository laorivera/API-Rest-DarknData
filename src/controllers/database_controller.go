package controllers

import (
	"builder/src/database"

	"github.com/gin-gonic/gin"
)

func DBtest(c *gin.Context) {
	var result int
	err := database.DB.QueryRow("SELECT 1").Scan(&result)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"db_test": result, "message": "Database is working!"})
}
