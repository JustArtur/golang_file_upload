package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"server/config"
	"strconv"
	"time"
)

func ValidateJWT(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Envs.JWTSecret), nil
	})
}

func GetUserIDFromJWT(token *jwt.Token) (int, error) {
	claims := token.Claims.(jwt.MapClaims)

	userIDStr, err := claims.GetSubject()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func NewJWT(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   strconv.Itoa(userID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(config.Envs.JWTExpiration))),
	})

	signedToken, err := token.SignedString([]byte(config.Envs.JWTSecret))

	if err != nil {
		log.Fatal(err)

		return "", err
	}

	return signedToken, nil
}
