package server

import "github.com/gin-gonic/gin"

type Server struct {
	server *gin.Engine
	port   string
}

func (s *Server) Start() error {
	return s.server.Run(s.port)
}

func (s *Server) CORS() {
	allowedOrigins := []string{
		"https://laorivera.github.io",
		"http://localhost:4200",
		"https://crm-resulting-toolbar-clay.trycloudflare.com",
		"*",
	}

	s.server.Use(func(c *gin.Context) {
		requestOrigin := c.Request.Header.Get("Origin")

		// Check if origin is allowed
		for i := 0; i < len(allowedOrigins); i++ {
			if allowedOrigins[i] == "*" || requestOrigin == allowedOrigins[i] {
				c.Writer.Header().Set("Access-Control-Allow-Origin", requestOrigin)
				break
			}
		}

		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})
}

func (s *Server) GetRouter() *gin.Engine {
	return s.server
}

func NewServer() *Server {
	router := gin.Default()
	server := &Server{server: router, port: ":8080"}
	server.CORS()
	return server
}
