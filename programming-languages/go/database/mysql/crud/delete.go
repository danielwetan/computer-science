package main

// Delete a user
func deleteUser(userID int64) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := db.Exec(query, userID)
	return err
}
