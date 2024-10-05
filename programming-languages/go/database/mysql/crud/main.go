package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID    int
	Name  string
	Email string
	Age   int
}

var db *sql.DB

func main() {
	// Establish connection to MySQL
	var err error
	dsn := "root:password@tcp(127.0.0.1:3306)/testdb"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if the connection is alive
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Example CRUD operations
	newUser := User{Name: "John Doe", Email: "john@example.com", Age: 30}

	// Create user
	insertedID, err := createUser(newUser)
	if err != nil {
		log.Fatalf("Error creating user: %v", err)
	}
	fmt.Printf("Inserted user with ID: %d\n", insertedID)

	// Read user
	user, err := getUser(insertedID)
	if err != nil {
		log.Fatalf("Error fetching user: %v", err)
	}
	fmt.Printf("Fetched user: %+v\n", user)

	// Update user
	user.Name = "John Smith"
	err = updateUser(*user)
	if err != nil {
		log.Fatalf("Error updating user: %v", err)
	}
	fmt.Println("Updated user.")

	// Delete user
	err = deleteUser(insertedID)
	if err != nil {
		log.Fatalf("Error deleting user: %v", err)
	}
	fmt.Println("Deleted user.")

	// Sample transaction
	err = sampleTransaction()
	if err != nil {
		log.Fatalf("Error during transaction: %v", err)
	}
	fmt.Println("Transaction completed.")
}
