package auth

import (
	"client/utils/request"
	"client/utils/response"
	"client/utils/token"
	"log"
	"net/http"
)

func Login(username, password string) {
	resp, err := http.DefaultClient.Do(request.CreateJSONRequest("POST", "/login", map[string]string{
		"username": username,
		"password": password,
	}))
	if err != nil {
		log.Fatalf("eror sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("error response from server: %v, %v", resp.Status, response.GetErrorFromResponse(resp))
	}

	responseData := response.UnmarshalResponse(resp)

	token.SaveToken(responseData["token"])

	log.Println("you have successfully logged in")
}

func Register(username, password string) {
	resp, err := http.DefaultClient.Do(request.CreateJSONRequest("POST", "/registration", map[string]string{
		"username": username,
		"password": password,
	}))
	if err != nil {
		log.Fatalf("sending request error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		log.Fatalf("server error response: %v, %v", resp.Status, response.GetErrorFromResponse(resp))
	}

	log.Println("User created")
}
