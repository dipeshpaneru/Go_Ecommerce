package utils

import (
	"net/http"
	"fmt"
	"encoding/json"
)

func ParseJson(r *http.Request, payload interface{}) error {
	if r.Body == nil {
		return fmt.Errorf("Request body is nil")
	}
	
	return json.NewDecoder(r.Body).Decode(payload)
}


func WriteJson(w http.ResponseWriter, statusCode int, payload interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	
	return json.NewEncoder(w).Encode(payload)
}

func WriteError(w http.ResponseWriter, statusCode int, err error) {
	errorResponse := map[string]string{"error": err.Error()}
	WriteJson(w, statusCode, errorResponse)
}