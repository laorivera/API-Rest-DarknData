package main

import (
	"builder/src/routes"
	"builder/src/server"
)

func main() {

	server := server.NewServer()

	router := server.GetRouter()

	routes.SetupRoutes(router)

	server.Start()

}
