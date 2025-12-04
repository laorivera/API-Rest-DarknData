package server

import (
	"builder/src/database"
	"builder/src/services"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	server *gin.Engine
	port   string
}

func (s *Server) Start() error {
	return s.server.Run(s.port) // 0.0.0.0 testing
}

func (s *Server) CORS() {

	//allowedOrigins := services.LoadConfig().AllowedOrigins

	s.server.Use(func(c *gin.Context) {
		//requestOrigin := c.Request.Header.Get("Origin")

		/*
			for i := 0; i < len(allowedOrigins); i++ {
				if requestOrigin == allowedOrigins[i] {
					c.Writer.Header().Set("Access-Control-Allow-Origin", requestOrigin)
					break
				}
			}
		*/
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // all origins
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-API-Key")
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

func (s *Server) AddLog() {
	s.server.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		log.Printf("METRIC path=%s method=%s status=%d size=%d duration=%v",
			c.Request.URL.Path,
			c.Request.Method,
			c.Writer.Status(),
			c.Writer.Size(),
			duration,
		)
	})
}
func (s *Server) InitDB() {
	err := database.Init()
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	log.Println("Database connected successfully")

}

func NewServer() *Server {
	router := gin.Default()
	server := &Server{server: router, port: services.LoadConfig().Port}
	server.CORS()
	server.AddLog()
	server.InitDB()

	return server
}
