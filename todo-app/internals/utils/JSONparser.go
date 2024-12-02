package utils

import (
	"encoding/json"
	"net/http"
)

func ParseJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func RespondJSON(w http.ResponseWriter, status int, data string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(data))
}

func RespondError(w http.ResponseWriter, status int, message string) {
	RespondJSON(w, status, message)
}
