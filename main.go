package main

import (
	"log"
	"net/http"
	"os"

	"chaos-engineering-go/chaos"
	"chaos-engineering-go/config"
	"chaos-engineering-go/handlers"
)

func main() {
	config.LoadEnv() // Load failure rates from environment

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.BasicHandler)

	// Wrap the router with Chaos middleware
	chaosMux := chaos.Middleware(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Starting server on port", port)
	err := http.ListenAndServe(":"+port, chaosMux)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
