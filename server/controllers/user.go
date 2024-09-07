package controllers

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"server/helpers"
	"server/models"
	"server/services/jwt"
	"server/types"
)

var user types.UserPayload

func Login(w http.ResponseWriter, r *http.Request) {
	log.Printf("Started %s %s", r.Method, r.RequestURI)
	defer log.Printf("Completed %s %s", r.Method, r.RequestURI)

	err := helpers.ParseRequest(r, &user)
	if err != nil {
		log.Println(err)
		helpers.SendResponse(w, http.StatusBadRequest, err)
		return
	}

	userRecord, err := models.GetUserByEmail(user.Email)
	if err != nil {
		log.Println(err)
		helpers.SendResponse(w, http.StatusBadRequest, fmt.Errorf("invalid user email or password"))
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userRecord.Password), []byte(user.Password)); err != nil {
		log.Println(err)
		helpers.SendResponse(w, http.StatusBadRequest, fmt.Errorf("invalid user email or password"))
		return
	}

	token, err := jwt.NewJWT(userRecord.ID)
	if err != nil {
		log.Println(err)
		helpers.SendResponse(w, http.StatusInternalServerError, err)
		return
	}

	helpers.SendResponse(w, http.StatusOK, map[string]string{"access_token": token})
}

func Register(w http.ResponseWriter, r *http.Request) {
	log.Printf("Started %s %s", r.Method, r.RequestURI)
	defer log.Printf("Completed %s %s", r.Method, r.RequestURI)

	err := helpers.ParseRequest(r, &user)
	if err != nil {
		helpers.SendResponse(w, http.StatusBadRequest, err)
		return
	}

	result, _ := models.GetUserByEmail(user.Email)
	if result != nil {
		helpers.SendResponse(w, http.StatusBadRequest, fmt.Errorf("user with this email already exists"))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, err)
		return
	}

	err = models.CreateUser(types.UserPayload{
		Email:    user.Email,
		Password: string(hashedPassword),
	})
	if err != nil {
		helpers.SendResponse(w, http.StatusInternalServerError, err)
		return
	}

	helpers.SendResponse(w, http.StatusCreated, nil)
}
