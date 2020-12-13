package main

import (
	"github.com/thyagofr/coodesh/desafio/http/api"
	"log"
	"net/http"
)

func main() {
	server := api.Routes()
	// utils.InitializeCron()
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
