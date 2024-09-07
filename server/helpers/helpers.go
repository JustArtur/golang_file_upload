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

	var response any

	if msg, ok := payload.(string); ok {
		response = map[string]string{"message": msg}
	} else {
		response = payload
	}

	json.NewEncoder(w).Encode(response)
}

func GetTokenFromRequest(r *http.Request) string {
	return r.Header.Get("Authorization")
}

func GetUserIDFromContext(r *http.Request) int {
	return r.Context().Value("userID").(int)
}
