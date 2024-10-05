package main

// Update a user
func updateUser(user User) error {
	query := "UPDATE users SET name = ?, email = ?, age = ? WHERE id = ?"
	_, err := db.Exec(query, user.Name, user.Email, user.Age, user.ID)
	return err
}
