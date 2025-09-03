package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// CORS
	r.Use(func(c *gin.Context) {

		allowedOrigins := []string{
			//"https://laorivera.github.io",
			//"http://10.8.0.0/24",
			"http://localhost:8080",
		}

		requestOrigin := c.Request.Header.Get("Origin")

		for _, origin := range allowedOrigins {
			if requestOrigin == origin {
				c.Writer.Header().Set("Access-Control-Allow-Origin", requestOrigin)
				break
			}
		}

		// headers
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		c.Next()
	})

	setupRoutes(r)

	// Corre en puerto 8080
	r.Run()

}
