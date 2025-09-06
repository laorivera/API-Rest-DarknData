package main

import (
	"fmt"
)

/*
func main() {

	itemmanager := NewItemManager()
	fmt.Println(itemmanager.ArmorsBySlot("Head", "1")[0])
	fmt.Println(Items.ItemsArmor[0])
	r := gin.Default()

	// CORS
	r.Use(func(c *gin.Context) {

		allowedOrigins := []string{
			"https://laorivera.github.io", // Angular app
			"http://localhost:4200",
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
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		c.Next()
	})

	setupRoutes(r)

	// Corre en puerto 8080
	r.Run()

}
*/

func main() {

	Im := NewItemManager()
	fmt.Println(Im.ArmorsBySlot("Head", "1")[0].Name)
	listhead := Im.ArmorsBySlot("Head", "1")
	var list []string
	for i := 0; i < len(listhead); i++ {
		list = append(list, listhead[i].Name)
	}
	fmt.Println(Items.ItemsArmor[0])
	fmt.Println(list)
}
