package main

// Create a new user
func createUser(user User) (int64, error) {
	query := "INSERT INTO users (name, email, age) VALUES (?, ?, ?)"
	result, err := db.Exec(query, user.Name, user.Email, user.Age)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
