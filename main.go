package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// CORS
	r.Use(func(c *gin.Context) {

		allowedOrigins := []string{
			"https://laorivera.github.io", // Angular app
			"http://localhost:4200",
			"http://10.8.0.0/24",
		}

		// Get the request's Origin
		requestOrigin := c.Request.Header.Get("Origin")

		for _, origin := range allowedOrigins {
			if requestOrigin == origin {
				c.Writer.Header().Set("Access-Control-Allow-Origin", requestOrigin)
				break
			}
		}

		// Required headers
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		c.Next()
	})

	setupRoutes(r)

	// Corre en puerto 8080
	r.Run()

}
