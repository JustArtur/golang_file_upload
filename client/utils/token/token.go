package token

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type TokenData struct {
	AccessToken string `json:"token"`
}

func SaveToken(token string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("error getting user homedir: %v", err)
	}

	configDir := filepath.Join(homeDir, ".golang-cli")
	os.MkdirAll(configDir, os.ModePerm)

	tokenFilePath := filepath.Join(configDir, "token.json")

	file, err := os.Create(tokenFilePath)
	if err != nil {
		log.Fatalf("error creating token file: %v", err)
	}
	defer file.Close()

	json.NewEncoder(file).Encode(TokenData{AccessToken: token})
}

func LoadToken() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("error getting user homedir: %v", err)
	}

	tokenFilePath := filepath.Join(homeDir, ".golang-cli", "token.json")
	file, err := os.Open(tokenFilePath)
	if err != nil {
		log.Fatalf("error opening token file: %v", err)
	}
	defer file.Close()

	var tokenData TokenData
	if err := json.NewDecoder(file).Decode(&tokenData); err != nil {
	}

	return tokenData.AccessToken
}
