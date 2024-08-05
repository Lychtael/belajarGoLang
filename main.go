package main

import (
	"belajarGoLang/config"
	"belajarGoLang/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config.ConnectDatabase()

	r := routes.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
