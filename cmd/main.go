package main

import (
	"log"
	"github.com/dipeshpaneru/Go_Ecommerce/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080", nil) // Replace nil with actual DB connection

	if err := server.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}