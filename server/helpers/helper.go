package helpers

import (
	"encoding/json"
	"net/http"
)

func ParseRequest(r *http.Request, payload any) error {
	return json.NewDecoder(r.Body).Decode(payload)
}

func SendResponse(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(payload)
}

func GetUserIDFromContext(r *http.Request) int {
	return r.Context().Value("userID").(int)
}
