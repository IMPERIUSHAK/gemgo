package handlers

import (
	"context"
	"encoding/json"
	"gemgo/gemini"
	"gemgo/models"
	"net/http"
	"time"
)

func MakeAskGemHandler(gClient *gemini.Client) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req models.Recipe

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		bytes, err := json.Marshal(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
		defer cancel()

		response, err := gClient.Generate(ctx, string(bytes), "gemini-3.5-flash-lite")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)
		if err!=nil{
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return 
		}
	}
}
