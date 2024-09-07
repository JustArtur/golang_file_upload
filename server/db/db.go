package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"server/config"
)

var Db *sql.DB

func ConnectToDB() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Envs.DbHost, config.Envs.DbPort, config.Envs.DbUser, config.Envs.DbPass, config.Envs.DbName, config.Envs.DbSSLMode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("database connect failed: %v", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("database ping failed: %v", err)
	}

	Db = db

	var dbName string
	err = db.QueryRow("SELECT current_database()").Scan(&dbName)
	if err != nil {
		log.Fatal("Error querying current database name: ", err)
	}

	log.Println("Successfully connected to database: ", dbName)
}
