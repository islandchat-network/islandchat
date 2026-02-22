package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/islandchat-network/islandchat/backend/internal/models"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	res := models.APIResponse{
		Status:  "ok",
		Message: "IslandChat Backend is running smoothly 🌴",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
