package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

// reference
// https://go.dev/doc/tutorial/database-access
func main() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		AllowNativePasswords: true,
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "first",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	db.SetMaxOpenConns(10)
	fmt.Println("open connection: ", db.Stats().OpenConnections)
	fmt.Println("max open connection: ", db.Stats().MaxOpenConnections)
}
