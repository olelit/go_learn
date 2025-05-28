package handlers

import (
	"encoding/json"
	"learn/go/internal/api/middleware"
	"learn/go/internal/requests"
	"learn/go/internal/services"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var req requests.RegisterRequest

	w.Header().Set("Content-Type", "application/json")

	if status, errMap := middleware.ValidateRequest(r, &req); status != 0 {
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(errMap)
		return
	}

	user, err := services.Register(req)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(user)
}
