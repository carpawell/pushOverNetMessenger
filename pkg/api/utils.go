package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// Function for responding when errors appear
func errorResponse(w http.ResponseWriter, err error, StatusCode int) {
	w.WriteHeader(StatusCode)
	if err := json.NewEncoder(w).Encode(Error{ErrorMessage: err.Error()}); err != nil {
		log.Fatalf("encoding:%s\n", err)
	}
}
