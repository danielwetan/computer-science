package main

import (
	"database/sql"
	"fmt"
)

// Read (get) a user by ID
func getUser(userID int64) (*User, error) {
	query := "SELECT id, name, email, age FROM users WHERE id = ?"
	row := db.QueryRow(query, userID)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User not found")
		}
		return nil, err
	}
	return &user, nil
}
