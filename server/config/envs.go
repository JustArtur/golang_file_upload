package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	JWTSecret     string
	JWTExpiration int64
	DbName        string
	DbHost        string
	DbPort        string
	DbUser        string
	DbPass        string
	DbSSLMode     string
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
		JWTSecret:     os.Getenv("JWT_SECRET"),
		JWTExpiration: getAsInt("JWT_EXPIRATION"),
		DbName:        os.Getenv("DB_NAME"),
		DbHost:        os.Getenv("DB_HOST"),
		DbPort:        os.Getenv("DB_PORT"),
		DbUser:        os.Getenv("DB_USER"),
		DbPass:        os.Getenv("DB_PASS"),
		DbSSLMode:     os.Getenv("DB_SSL_MODE"),
	}
}

func getAsInt(key string) int64 {
	env, err := strconv.ParseInt(os.Getenv(key), 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	return env
}
