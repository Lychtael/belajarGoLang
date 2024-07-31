package main

import (
	"belajarGoLang/routes"
)

func main() {
	router := routes.SetupRouter()
	router.Run(":8080")
}
