package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	ServerHost string
}

var Envs Config

func InitEnvs() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	err := godotenv.Load(".env." + env)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("env: succesfull loaded file: %s", ".env."+env)
	}

	Envs = loadEnvs()
}

func loadEnvs() Config {
	return Config{
		ServerHost: os.Getenv("SERVER_HOST"),
	}
}
