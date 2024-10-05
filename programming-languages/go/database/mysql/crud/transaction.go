package main

// Sample transaction with two operations
func sampleTransaction() error {
	// Begin a new transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Defer a rollback in case something fails
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Transactional step 1: Insert a new user
	insertUserQuery := "INSERT INTO users (name, email, age) VALUES (?, ?, ?)"
	result, err := tx.Exec(insertUserQuery, "Alice Doe", "alice@example.com", 25)
	if err != nil {
		return err
	}
	newUserID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	// Transactional step 2: Update the user's email
	updateUserQuery := "UPDATE users SET email = ? WHERE id = ?"
	_, err = tx.Exec(updateUserQuery, "alice.updated@example.com", newUserID)
	if err != nil {
		return err
	}

	// Commit the transaction
	return tx.Commit()
}
