package handlers

import (
	"encoding/json"
	"learn/go/internal/modes"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var user modes.Users
	json.NewDecoder(r.Body).Decode(&user)

	modes.Register(user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(modes.List)
}
