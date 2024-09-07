package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"server/helpers"
	"server/models"
	"server/services/jwt"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := helpers.GetTokenFromRequest(r)

		token, err := jwt.ValidateJWT(tokenStr)
		if err != nil {
			log.Printf("validate token error: %v", err)
			unauthorized(w)
			return
		}

		userID, err := jwt.GetUserIDFromJWT(token)
		if err != nil {
			log.Printf("failed to get user id from token: %v", err)
			unauthorized(w)
			return
		}

		u, err := models.GetUserByID(userID)
		if err != nil {
			log.Printf("failed to get user by id: %v", err)
			unauthorized(w)
			return
		}

		ctx := context.WithValue(r.Context(), "userID", u.ID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func unauthorized(w http.ResponseWriter) {
	helpers.SendResponse(w, http.StatusUnauthorized, fmt.Errorf("unauthorized"))
}
