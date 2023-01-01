package main

import (
	"fmt"
	"log"
	"os/user"
)

// Resolving the user home directory

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The username: " + usr.Username)
	fmt.Println("The user home directory: " + usr.HomeDir)
}
