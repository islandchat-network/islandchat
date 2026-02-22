package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/islandchat-network/islandchat/backend/internal/api/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handlers.HealthCheck)

	port := ":8080"
	fmt.Printf("🚀 IslandChat Server started at http://localhost%s\n", port)

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Critical error starting server: %v", err)
	}
}
