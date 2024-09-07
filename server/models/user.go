package models

import (
	"fmt"
	"log"
	"server/db"
	"server/types"
)

func GetUserByEmail(email string) (*types.UserPayload, error) {
	user := new(types.UserPayload)

	query := "SELECT * FROM users WHERE \"email\" = $1"

	log.Print("pq: ", query, email)
	rows, err := db.Db.Query(query, email)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Email,
			&user.Password,
		)
		if err != nil {
			return nil, err
		}
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found, incorrect email or password")
	}

	return user, nil
}

func GetUserByID(ID int) (*types.UserPayload, error) {
	user := new(types.UserPayload)

	query := "SELECT * FROM users WHERE \"id\" = $1"

	log.Print("pq: ", query, ID)
	rows, err := db.Db.Query(query, ID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Email,
			&user.Password,
		)
		if err != nil {
			return nil, err
		}
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func CreateUser(user types.UserPayload) error {

	query := "INSERT INTO users (email, password) VALUES ($1, $2)"
	log.Print("pq: ", query, user.Email, user.Password)
	_, err := db.Db.Exec(query, user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil
}
