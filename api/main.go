package main

import (
	"log"
	"net/http"
	"server/routes"
)

func main() {
	router := routes.SetupRouter()
	log.Println("APP started and running on PORT 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Error starting server", err)
	}
}
