// main entry point for our server.

package main

import (
	"backend/models"
	"fmt"
)

func main() {
	fmt.Println("Hello. More work to be done.")
	user := models.User{
		ID:       "1",
		Name:     "Bob",
		Age:      60,
		Location: "North Ikea",
	}
	fmt.Println("More work to be done. again...")
	fmt.Printf("Yo: %v\n", user)
}
