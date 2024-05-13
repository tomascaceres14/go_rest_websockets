package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/tomascaceres14/rest_websockets/server"
)

type JsonResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(JsonResponse{
			Message: "API Working",
			Status:  true,
		})

	}
}
