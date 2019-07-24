package utils

import (
	"encoding/json"
	"net/http"
)

// Message creates a message wrapper to send over http
func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

// Respond sends a message over http
func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Contains returns true if element exist in the array
func Contains(elem int32, elements []int32) bool {
	for _, e := range elements {
		if e == elem {
			return true
		}
	}

	return false
}
