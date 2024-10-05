package main

import (
	"fmt"
)

// Reference
// https://go.dev/talks/2013/go4python.slide#47
func main() {
	fmt.Println("before monkey patching...")
	sayHi("daniel")
	sayHi("admin")

	fmt.Println("after monkey patching...")
	auth = func(user string) bool { return true }
	sayHi("daniel")

	auth = func(user string) bool { return false }
	sayHi("admin")
}

func sayHi(user string) {
	if !auth(user) {
		fmt.Printf("unknown user %v\n", user)
		return
	}
	fmt.Printf("Hi, %v\n", user)
}

var auth = func(user string) bool {
	return user == "admin"
}
