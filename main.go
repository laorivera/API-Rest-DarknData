package main

import (
	"builder/src/server"
)

func main() {

	server := server.NewServer()

	router := server.GetRouter()

	setupRoutes(router)

	// Corre en puerto 8080
	server.Start()

}
