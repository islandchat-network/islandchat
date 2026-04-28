package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/islandchat-network/islandchat/backend/internal/api/handlers"
	"github.com/islandchat-network/islandchat/backend/internal/api/ws"
)

func main() {
	// Initialize the Hub
	hub := ws.NewHub()
	go hub.Run()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", handlers.HealthCheck)

	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWs(hub, w, r)
	})

	port := ":8080"
	fmt.Printf("🚀 IslandChat Server started at http://localhost%s\n", port)

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Critical error starting server: %v", err)
	}
}
